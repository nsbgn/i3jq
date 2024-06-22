module {
  name: "ipc",
  description: "Filters for inter-process communication."
};

def run_command(payload): _i3jq(0; payload);
def get_workspaces: _i3jq(1);
def subscribe(payload): _i3jq(2; payload | tostring; true);
def get_outputs: _i3jq(3);
def get_tree: _i3jq(4);
def get_marks: _i3jq(5);
def get_bar_config(payload): _i3jq(6; payload);
def get_bar_config: get_bar_config("");
def get_version: _i3jq(7);
def get_binding_modes: _i3jq(8);
def get_config: _i3jq(9);
def send_tick: _i3jq(10);
def sync(payload): _i3jq(11; payload);
def get_binding_state(payload): _i3jq(12; payload);
