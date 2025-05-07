# This is a long-running script that implements a master-stack layout. That 
# means that there is one "main" window and any additional window gets stacked 
# to the side.

# Note: To normalize a workspace:
# 1. Put workspace in the correct orientation
# 2. Identify the stack
# 3. Flatten the stack
# 4. Identify all windows that are *not* the stack.
# 5. Move those to the front/back of the stack too.
# 6. Move first/last window(s) in a direction orthogonal to the stack, until we
# are in n-mode.
# 7. Synchronize (ipc::send_tick("synchronize"))

import "builtin/ipc" as ipc;
import "builtin/tree" as tree;

# The mark to which to send new windows
def INSERT: "insert";
# The mark with which to swap new windows
def SWAP: "swap";

def do(commands):
  [commands] |
  flatten |
  if . == [] then
    empty
  else
    join("; ") |
    { command: .,
      result: ipc::run_command(.) }
  end;

def is_event(event; change):
  (.event as $e | any(event; $e == .)) and
  (.change as $c | any(change; $c == .));

def mark($mark; $yes):
  if $yes then
    if .marks | any(. == $mark) | not then
      "[con_id=\(.id)] mark --add \($mark)"
    else
      empty
    end
  else
    "unmark \($mark)"
  end;
def mark($mark):
  mark($mark; true);

def position:
  if . == "top" or . == "left" then 0 # beginning
  elif . == "bottom" or . == "right" then -1 # end
  else "Unknown position '\(.)'" | error end;

def axis:
  if . == "left" or . == "right" then 0 # horizontal axis
  elif . == "bottom" or . == "top" then 1 # vertical axis
  else "Unknown axis '\(.)'" | error end;

# Is the container in the correct layout?
def is_layout($orientation):
  .nodes |
  length == 0 or (length == 1 and .[0] | (
    ["splith", "splitv"] as $split |
    ($orientation | axis) as $axis |
    ($orientation | position) as $i |
    .layout == $split[$axis]
    and .nodes[$i].layout == $split[1-$axis] 
    and all(.nodes[:$i], .nodes[$i+1:], .nodes[$i].nodes[]; .layout == "none")
  ));

# Ensure that all marks are in the correct spot. We can assume that the layout
# is correct here; if it isn't, it will be fixed later.
def ensure_marks($orientation):
  ($orientation | position) as $i |
  (.nodes[0] // empty) |
  (.nodes | length < 2) as $monocle |
  (.nodes[$i] // empty) |
  tree::focused |
  mark(INSERT),
  mark(SWAP; $monocle);

# Organize a workspace into a master-stack layout. The master-stack layout is a
# single container on the top level (the layout-root) containing one leaf
# window (the master) and potentially one more container (the stack). The
# latter contains the rest of the leaf windows. This filter makes no assumption
# about the current state of the layout tree.
def apply_layout($orientation):
  (.nodes | length) as $n |
  ["splith", "splitv"] as $split |
  ["horizontal", "vertical"] as $splits |
  ($orientation | axis) as $axis |
  ($orientation | position) as $i |

  # If the workspace is now entirely empty, we just need to make sure that any
  # new window opened won't appear in some other workspace.
  if $n == 0 then
    "unmark \(SWAP); unmark \(INSERT)"
  else
    # Otherwise, we can descend into the layout-root.
    # TODO: Move excess windows
    .nodes[0] |
    (.nodes | length) as $n |

    # If the layout-root is itself a leaf node, it needs to be split and
    # correct marks set
    if $n == 0 then
      mark(INSERT),
      mark(SWAP; $i == 0),
      "[con_id=\(.id)] split \($splits[$axis])"

    # If it is split already, but contains only one node:
    elif $n == 1 then
      .nodes[0] |
      (.nodes | length) as $n |

      if $n == 0 then
        mark(INSERT),
        mark(SWAP; $i == 0)
      elif $n == 1 then
        # TODO: We assume here that the node underneath *this* is a leaf, which
        # is not necessarily true
        .nodes[0] |
        "[con_id=\(.id)] split none",
        mark(INSERT),
        mark(SWAP; $i == 0)
      else
        # When there is only one top-level node, we want to assume that there
        # is also just one master window, but that might not be true if we just
        # closed the previous master window. Then this node is the stack. In
        # that case, we select the second most recently focused window in this
        # stack and promote it to master
        (tree::focused_child(1) | "[con_id=\(.id)] move right"), # TODO
        (tree::focused_child(0) |
        mark(INSERT),
        mark(SWAP; $i == 0))
      end

    # We already have both a master and stack container
    # TODO: We assume that we *only* have these
    elif $n > 1 then
      .nodes[-(1+$i)] as $master |
      .nodes[$i] as $stack |
      $stack |
      if .layout == "none" then
        mark(INSERT),
        "[con_id=\(.id)] split \($split[1-$axis])"
      else
        tree::focused |
        mark(INSERT)
      end,
      mark(SWAP; false)
    end
  end;

def init:
  # To instantly put new tiling windows where they belong, without a moment of 
  # flickering as the script responds to events, we start by putting in place 
  # rules to insert new windows after the window with the `INSERT` mark. To put 
  # it *before* that window, also set the `SWAP` mark.
  "for_window [tiling] move container to mark \(INSERT)",
  "for_window [tiling] swap container with mark \(SWAP)";

def main:
  do(init),
  foreach ipc::subscribe(["workspace", "window", "tick"]) as $e (
    {stack: "left"};
    if $e.event == "tick" then
      {stack: .stack, payload: $e.payload}
    end;
    . as {$stack, $n} | $e |
    if is_event("window"; "new", "close") or is_event("workspace"; "focus") then
      do(
        ipc::get_tree |
        tree::focused(.type == "workspace") |
        apply_layout($stack)
      )
    elif is_event("window"; "focus") then
      do(
        ipc::get_tree |
        tree::focused(.type == "workspace") |
        ensure_marks($stack)
      )
    else
      empty
    end
  );

main
