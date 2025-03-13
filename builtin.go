// Code generated by go generate; DO NOT EDIT.
// See builtin.generator.go for more information
package main
import . "github.com/itchyny/gojq"
import "github.com/itchyny/gojq"

func LoadBuiltin(name string) *gojq.Query {
	switch(name){
	case "ansi":
		return &Query{Meta: &ConstObject{KeyVals: []*ConstObjectKeyVal{{Key: "name", Val: &ConstTerm{Str: "ansi"}}, {Key: "description", Val: &ConstTerm{Str: "ANSI escape codes."}}}}, FuncDefs: []*FuncDef{{Name: "color", Args: []string{"$c", "$fg"}, Body: &Query{FuncDefs: []*FuncDef{{Name: "_colors", Body: &Query{Term: &Term{Type: gojq.TermTypeObject, Object: &Object{KeyVals: []*ObjectKeyVal{{KeyString: &String{Str: "black"}, Val: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "30"}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "40"}}}}}}}, {KeyString: &String{Str: "red"}, Val: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "31"}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "41"}}}}}}}, {KeyString: &String{Str: "green"}, Val: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "32"}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "42"}}}}}}}, {KeyString: &String{Str: "yellow"}, Val: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "33"}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "43"}}}}}}}, {KeyString: &String{Str: "blue"}, Val: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "34"}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "44"}}}}}}}, {KeyString: &String{Str: "magenta"}, Val: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "35"}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "45"}}}}}}}, {KeyString: &String{Str: "cyan"}, Val: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "36"}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "46"}}}}}}}, {KeyString: &String{Str: "white"}, Val: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "37"}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "47"}}}}}}}, {KeyString: &String{Str: "lightgray"}, Val: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "37"}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "47"}}}}}}}, {KeyString: &String{Str: "gray"}, Val: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "90"}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "100"}}}}}}}, {KeyString: &String{Str: "brightblack"}, Val: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "90"}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "100"}}}}}}}, {KeyString: &String{Str: "brightred"}, Val: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "91"}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "101"}}}}}}}, {KeyString: &String{Str: "brightgreen"}, Val: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "92"}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "102"}}}}}}}, {KeyString: &String{Str: "brightyellow"}, Val: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "93"}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "103"}}}}}}}, {KeyString: &String{Str: "brightblue"}, Val: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "94"}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "104"}}}}}}}, {KeyString: &String{Str: "brightmagenta"}, Val: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "95"}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "105"}}}}}}}, {KeyString: &String{Str: "brightcyan"}, Val: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "96"}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "106"}}}}}}}, {KeyString: &String{Str: "brightwhite"}, Val: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "97"}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "107"}}}}}}}}}}}}}, Term: &Term{Type: gojq.TermTypeString, Str: &String{Queries: []*Query{{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "\x1b["}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_colors"}, SuffixList: []*Suffix{{Index: &Index{Start: &Query{Func: "$c"}}}, {Index: &Index{Start: &Query{Term: &Term{Type: gojq.TermTypeIf, If: &If{Cond: &Query{Func: "$fg"}, Then: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "0"}}, Else: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "1"}}}}}}}}}}}}, {Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "m"}}}}}}}}, {Name: "_wrap", Args: []string{"$i"}, Body: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Queries: []*Query{{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "\x1b["}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Func: "$i"}}}, {Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "m"}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Func: "."}}}, {Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "\x1b[0m"}}}}}}}}, {Name: "bold", Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_wrap", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "1"}}}}}}}, {Name: "dim", Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_wrap", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "2"}}}}}}}, {Name: "italic", Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_wrap", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "3"}}}}}}}, {Name: "underline", Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_wrap", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "4"}}}}}}}, {Name: "invert", Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_wrap", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "7"}}}}}}}, {Name: "clear", Body: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "\x1b[2J"}}}}, {Name: "curpos", Args: []string{"$i", "$j"}, Body: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Queries: []*Query{{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "\x1b["}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Func: "$i"}}}, {Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: ";"}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Func: "$j"}}}, {Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "H"}}}}}}}}, {Name: "fg", Args: []string{"$c"}, Body: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Queries: []*Query{{Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "color", Args: []*Query{{Func: "$c"}, {Func: "true"}}}}}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Func: "."}}}, {Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "\x1b[39m"}}}}}}}}, {Name: "bg", Args: []string{"$c"}, Body: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Queries: []*Query{{Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "color", Args: []*Query{{Func: "$c"}, {Func: "false"}}}}}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Func: "."}}}, {Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "\x1b[49m"}}}}}}}}}}
	case "ipc":
		return &Query{Meta: &ConstObject{KeyVals: []*ConstObjectKeyVal{{Key: "name", Val: &ConstTerm{Str: "ipc"}}, {Key: "description", Val: &ConstTerm{Str: "Filters for inter-process communication."}}}}, FuncDefs: []*FuncDef{{Name: "run_command", Args: []string{"$payload"}, Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_internal", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "0"}}, {Func: "$payload"}, {Func: "false"}}}}}}, {Name: "get_workspaces", Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_internal", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "1"}}, {Func: "null"}, {Func: "false"}}}}}}, {Name: "subscribe", Args: []string{"$payload"}, Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_internal", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "2"}}, {Left: &Query{Func: "$payload"}, Op: gojq.OpPipe, Right: &Query{Func: "tostring"}}, {Func: "true"}}}}}}, {Name: "get_outputs", Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_internal", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "3"}}, {Func: "null"}, {Func: "false"}}}}}}, {Name: "get_tree", Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_internal", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "4"}}, {Func: "null"}, {Func: "false"}}}}}}, {Name: "get_marks", Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_internal", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "5"}}, {Func: "null"}, {Func: "false"}}}}}}, {Name: "get_bar_config", Args: []string{"$payload"}, Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_internal", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "6"}}, {Func: "$payload"}, {Func: "false"}}}}}}, {Name: "get_bar_config", Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "get_bar_config", Args: []*Query{{Func: "null"}}}}}}, {Name: "get_version", Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_internal", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "7"}}, {Func: "null"}, {Func: "false"}}}}}}, {Name: "get_binding_modes", Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_internal", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "8"}}, {Func: "null"}, {Func: "false"}}}}}}, {Name: "get_config", Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_internal", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "9"}}, {Func: "null"}, {Func: "false"}}}}}}, {Name: "send_tick", Args: []string{"$payload"}, Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_internal", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "10"}}, {Func: "$payload"}, {Func: "false"}}}}}}, {Name: "sync", Args: []string{"$payload"}, Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_internal", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "11"}}, {Func: "$payload"}, {Func: "false"}}}}}}, {Name: "get_binding_state", Args: []string{"$payload"}, Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "_internal", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "12"}}, {Func: "$payload"}, {Func: "false"}}}}}}, {Name: "do", Args: []string{"commands"}, Body: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Func: "commands"}}}}, Op: gojq.OpPipe, Right: &Query{Left: &Query{Func: "flatten"}, Op: gojq.OpPipe, Right: &Query{Term: &Term{Type: gojq.TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "."}, Op: gojq.OpEq, Right: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{}}}}, Then: &Query{Func: "empty"}, Else: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "join", Args: []*Query{{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "; "}}}}}}}, Op: gojq.OpPipe, Right: &Query{Term: &Term{Type: gojq.TermTypeObject, Object: &Object{KeyVals: []*ObjectKeyVal{{Key: "command", Val: &Query{Func: "."}}, {Key: "result", Val: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "run_command", Args: []*Query{{Func: "."}}}}}}}}}}}}}}}}}}}
	case "show":
		return &Query{Meta: &ConstObject{KeyVals: []*ConstObjectKeyVal{{Key: "name", Val: &ConstTerm{Str: "show"}}, {Key: "description", Val: &ConstTerm{Str: "A module to show a readable tiling tree. This is the default module."}}}}, Imports: []*Import{{ImportPath: "builtin/ipc", ImportAlias: "ipc"}, {ImportPath: "builtin/tree", ImportAlias: "tree"}, {ImportPath: "builtin/ansi", ImportAlias: "ansi"}}, FuncDefs: []*FuncDef{{Name: "hex", Body: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Left: &Query{Left: &Query{Func: "."}, Op: gojq.OpDiv, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "16"}}}, Op: gojq.OpPipe, Right: &Query{Left: &Query{Func: "floor"}, Op: gojq.OpPipe, Right: &Query{Term: &Term{Type: gojq.TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "."}, Op: gojq.OpGt, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "0"}}}, Then: &Query{Func: "hex"}, Else: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{}}}}}}}}}}, Op: gojq.OpAdd, Right: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "0123456789abcdef"}, SuffixList: []*Suffix{{Index: &Index{Start: &Query{Left: &Query{Func: "."}, Op: gojq.OpMod, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "16"}}}}}}}}}}, {Name: "pad", Args: []string{"$n"}, Body: &Query{Left: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: " "}}}, Op: gojq.OpMul, Right: &Query{Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Left: &Query{Func: "$n"}, Op: gojq.OpSub, Right: &Query{Func: "length"}}}}}, Op: gojq.OpAdd, Right: &Query{Func: "."}}}, {Name: "truncate", Args: []string{"$n"}, Body: &Query{Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Left: &Query{Left: &Query{Func: "$n"}, Op: gojq.OpDiv, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "2"}}}, Op: gojq.OpPipe, Right: &Query{Func: "floor"}}, SuffixList: []*Suffix{{Bind: &Bind{Patterns: []*Pattern{{Name: "$m"}}, Body: &Query{Term: &Term{Type: gojq.TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "length"}, Op: gojq.OpGt, Right: &Query{Func: "$n"}}, Then: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Queries: []*Query{{Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Start: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "0"}}, End: &Query{Left: &Query{Func: "$m"}, Op: gojq.OpSub, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "2"}}}, IsSlice: true}}}}}, {Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "(…)"}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Start: &Query{Left: &Query{Left: &Query{Func: "length"}, Op: gojq.OpSub, Right: &Query{Func: "$m"}}, Op: gojq.OpAdd, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "1"}}}, End: &Query{Func: "length"}, IsSlice: true}}}}}}}}}}}}}}}}}}, {Name: "show", Args: []string{"head", "tail"}, Body: &Query{FuncDefs: []*FuncDef{{Name: "layout", Body: &Query{Term: &Term{Type: gojq.TermTypeObject, Object: &Object{KeyVals: []*ObjectKeyVal{{Key: "splith", Val: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "H"}}}}, {Key: "splitv", Val: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "V"}}}}, {Key: "tabbed", Val: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "T"}}}}, {KeyString: &String{Str: "stacked"}, Val: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "S"}}}}}}, SuffixList: []*Suffix{{Index: &Index{Start: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "layout"}}}}}}}}}, {Name: "node", Body: &Query{Term: &Term{Type: gojq.TermTypeIf, If: &If{Cond: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "type"}}}, Op: gojq.OpEq, Right: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "root"}}}}, Then: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: " - "}}}, Elif: []*IfElif{{Cond: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "type"}}}, Op: gojq.OpEq, Right: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "output"}}}}, Then: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: " o "}}}}, {Cond: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "type"}}}, Op: gojq.OpEq, Right: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "workspace"}}}}, Then: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Queries: []*Query{{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: " "}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Func: "layout"}}}, {Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: " "}}}}}}}}, {Cond: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "layout"}}}, Op: gojq.OpNe, Right: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "none"}}}}, Then: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Queries: []*Query{{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "·"}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Func: "layout"}}}, {Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "·"}}}}}}}}}, Else: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{}}}}}}}, {Name: "show_aux", Args: []string{"$prefix", "$prefix_child", "$prefix_parent", "$on_focus_path"}, Body: &Query{Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "tree::focus_child"}, SuffixList: []*Suffix{{Index: &Index{Name: "id"}}}}}, Op: gojq.OpAlt, Right: &Query{Func: "null"}}, SuffixList: []*Suffix{{Bind: &Bind{Patterns: []*Pattern{{Name: "$focus_id"}}, Body: &Query{Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "floating_nodes"}, SuffixList: []*Suffix{{Index: &Index{Start: &Query{Term: &Term{Type: gojq.TermTypeUnary, Unary: &Unary{Op: gojq.OpSub, Term: &Term{Type: gojq.TermTypeNumber, Number: "1"}}}}}}}}}, Op: gojq.OpAlt, Right: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "nodes"}, SuffixList: []*Suffix{{Index: &Index{Start: &Query{Term: &Term{Type: gojq.TermTypeUnary, Unary: &Unary{Op: gojq.OpSub, Term: &Term{Type: gojq.TermTypeNumber, Number: "1"}}}}}}}}}}, SuffixList: []*Suffix{{Index: &Index{Name: "id"}}, {Bind: &Bind{Patterns: []*Pattern{{Name: "$last_id"}}, Body: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Queries: []*Query{{Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Func: "head"}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Func: "$prefix"}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Func: "$prefix_parent"}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Left: &Query{Func: "node"}, Op: gojq.OpPipe, Right: &Query{Func: "ansi::invert"}}}}, {Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: " "}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Left: &Query{Func: "tail"}, Op: gojq.OpAlt, Right: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{}}}}}}}}}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeForeach, Foreach: &Foreach{Query: &Query{Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "nodes"}, SuffixList: []*Suffix{{Iter: true}}}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "floating_nodes"}, SuffixList: []*Suffix{{Iter: true}}}}}}}, Pattern: &Pattern{Name: "$node"}, Start: &Query{Func: "$on_focus_path"}, Update: &Query{Left: &Query{Func: "."}, Op: gojq.OpAnd, Right: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "$node"}, SuffixList: []*Suffix{{Index: &Index{Name: "id"}}}}}, Op: gojq.OpNe, Right: &Query{Func: "$focus_id"}}}, Extract: &Query{Term: &Term{Type: gojq.TermTypeIdentity, SuffixList: []*Suffix{{Bind: &Bind{Patterns: []*Pattern{{Name: "$prefocus"}}, Body: &Query{Left: &Query{Func: "$node"}, Op: gojq.OpPipe, Right: &Query{Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Left: &Query{Func: "$on_focus_path"}, Op: gojq.OpAnd, Right: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "id"}}}, Op: gojq.OpEq, Right: &Query{Func: "$focus_id"}}}, SuffixList: []*Suffix{{Bind: &Bind{Patterns: []*Pattern{{Name: "$focus"}}, Body: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIf, If: &If{Cond: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "id"}}}, Op: gojq.OpNe, Right: &Query{Func: "$last_id"}}, Then: &Query{Term: &Term{Type: gojq.TermTypeIf, If: &If{Cond: &Query{Func: "$focus"}, Then: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "│"}}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "┡"}}}}, Elif: []*IfElif{{Cond: &Query{Func: "$prefocus"}, Then: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "┃"}}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "┠"}}}}}}, Else: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "│"}}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "├"}}}}}}}, Else: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: " "}}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeIf, If: &If{Cond: &Query{Func: "$focus"}, Then: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "┗"}}}, Else: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "└"}}}}}}}}}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeIf, If: &If{Cond: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "type"}}}, Op: gojq.OpEq, Right: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "floating_con"}}}}, Then: &Query{Term: &Term{Type: gojq.TermTypeIf, If: &If{Cond: &Query{Func: "$focus"}, Then: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "┅┅"}}}, Else: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "┄┄"}}}}}}, Else: &Query{Term: &Term{Type: gojq.TermTypeIf, If: &If{Cond: &Query{Func: "$focus"}, Then: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "━━"}}}, Else: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "──"}}}}}}}}}}}, SuffixList: []*Suffix{{Bind: &Bind{Patterns: []*Pattern{{Array: []*Pattern{{Name: "$x"}, {Name: "$y"}, {Name: "$z"}}}}, Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "show_aux", Args: []*Query{{Left: &Query{Func: "$prefix"}, Op: gojq.OpAdd, Right: &Query{Func: "$prefix_child"}}, {Term: &Term{Type: gojq.TermTypeString, Str: &String{Queries: []*Query{{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: " "}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Func: "$x"}}}, {Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "  "}}}}}}}, {Term: &Term{Type: gojq.TermTypeString, Str: &String{Queries: []*Query{{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: " "}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Func: "$y"}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Func: "$z"}}}}}}}, {Func: "$focus"}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}, Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "show_aux", Args: []*Query{{Term: &Term{Type: gojq.TermTypeString, Str: &String{}}}, {Term: &Term{Type: gojq.TermTypeString, Str: &String{}}}, {Term: &Term{Type: gojq.TermTypeString, Str: &String{}}}, {Func: "true"}}}}}}, {Name: "show", Args: []string{"tail"}, Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "show", Args: []*Query{{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: " "}}}, {Func: "tail"}}}}}}, {Name: "show", Body: &Query{FuncDefs: []*FuncDef{{Name: "head", Body: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "id"}}}, Op: gojq.OpPipe, Right: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIf, If: &If{Cond: &Query{Left: &Query{Func: "."}, Op: gojq.OpLt, Right: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "2147483646"}}}, Then: &Query{Func: "tostring"}, Else: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "·"}}}}}}, Op: gojq.OpPipe, Right: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "pad", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "5"}}}}}}, Op: gojq.OpAdd, Right: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: " "}}}}}}}, {Name: "tail", Body: &Query{Term: &Term{Type: gojq.TermTypeIf, If: &If{Cond: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "type"}}}, Op: gojq.OpEq, Right: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "root"}}}}, Then: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{}}}, Elif: []*IfElif{{Cond: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "type"}}}, Op: gojq.OpEq, Right: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "output"}}}}, Then: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Queries: []*Query{{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "output "}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "name"}}}}}}}}}, Op: gojq.OpPipe, Right: &Query{Func: "ansi::bold"}}}, {Cond: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "type"}}}, Op: gojq.OpEq, Right: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "workspace"}}}}, Then: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Queries: []*Query{{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "workspace "}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "name"}}}}}}}}}, Op: gojq.OpPipe, Right: &Query{Func: "ansi::bold"}}}, {Cond: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "layout"}}}, Op: gojq.OpNe, Right: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "none"}}}}, Then: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "tile"}}}}}, Else: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Queries: []*Query{{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "["}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "app_id"}}}, Op: gojq.OpPipe, Right: &Query{Func: "ansi::italic"}}}}, {Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "] "}}}, {Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "name"}}}, Op: gojq.OpPipe, Right: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "truncate", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "30"}}}}}}}}}}}}}}}}}}, Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "show", Args: []*Query{{Func: "head"}, {Func: "tail"}}}}}}, {Name: "watch", Body: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "ipc::subscribe", Args: []*Query{{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "window"}}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "workspace"}}}}}}}}}}}, Op: gojq.OpPipe, Right: &Query{Left: &Query{Func: "ipc::get_tree"}, Op: gojq.OpPipe, Right: &Query{Left: &Query{Left: &Query{Func: "ansi::clear"}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "ansi::curpos", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "1"}}, {Term: &Term{Type: gojq.TermTypeNumber, Number: "1"}}}}}}}, Op: gojq.OpComma, Right: &Query{Func: "show"}}}}}, {Name: "watch", Args: []string{"tail"}, Body: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "ipc::subscribe", Args: []*Query{{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "window"}}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "workspace"}}}}}}}}}}}, Op: gojq.OpPipe, Right: &Query{Left: &Query{Func: "ipc::get_tree"}, Op: gojq.OpPipe, Right: &Query{Left: &Query{Left: &Query{Func: "ansi::clear"}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "ansi::curpos", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "1"}}, {Term: &Term{Type: gojq.TermTypeNumber, Number: "1"}}}}}}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "show", Args: []*Query{{Func: "tail"}}}}}}}}}}, Left: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "ipc::get_tree"}}}, Op: gojq.OpPipe, Right: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "show"}}}}
	case "tree":
		return &Query{Meta: &ConstObject{KeyVals: []*ConstObjectKeyVal{{Key: "name", Val: &ConstTerm{Str: "tree"}}, {Key: "description", Val: &ConstTerm{Str: "Filters for navigating the layout tree."}}}}, FuncDefs: []*FuncDef{{Name: "children", Body: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "nodes"}, SuffixList: []*Suffix{{Iter: true}}}}, Op: gojq.OpComma, Right: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "floating_nodes"}, SuffixList: []*Suffix{{Iter: true}}}}}}, {Name: "focus_child", Args: []string{"$n"}, Body: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "focus"}, SuffixList: []*Suffix{{Index: &Index{Start: &Query{Func: "$n"}}}, {Bind: &Bind{Patterns: []*Pattern{{Name: "$id"}}, Body: &Query{Left: &Query{Func: "children"}, Op: gojq.OpPipe, Right: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "id"}}}, Op: gojq.OpEq, Right: &Query{Func: "$id"}}}}}}}}}}}}}, {Name: "focus_child", Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "focus_child", Args: []*Query{{Term: &Term{Type: gojq.TermTypeNumber, Number: "0"}}}}}}}, {Name: "focus", Args: []string{"condition"}, Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "until", Args: []*Query{{Func: "condition"}, {Func: "focus_child"}}}}}}, {Name: "focus", Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "focus", Args: []*Query{{Left: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "nodes"}}}, Op: gojq.OpEq, Right: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{}}}}, Op: gojq.OpAnd, Right: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "floating_nodes"}}}, Op: gojq.OpEq, Right: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{}}}}}}}}}}, {Name: "find_all", Args: []string{"condition"}, Body: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "recurse", Args: []*Query{{Func: "children"}}}}}, Op: gojq.OpPipe, Right: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Func: "condition"}}}}}}}, {Name: "find", Args: []string{"condition"}, Body: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "first", Args: []*Query{{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "find_all", Args: []*Query{{Func: "condition"}}}}}}}}}, Op: gojq.OpAlt, Right: &Query{Func: "null"}}}, {Name: "lineage", Args: []string{"target", "child"}, Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "target"}, SuffixList: []*Suffix{{Bind: &Bind{Patterns: []*Pattern{{Name: "$x"}}, Body: &Query{Term: &Term{Type: gojq.TermTypeIf, If: &If{Cond: &Query{Left: &Query{Left: &Query{Func: "$x"}, Op: gojq.OpEq, Right: &Query{Func: "true"}}, Op: gojq.OpOr, Right: &Query{Term: &Term{Type: gojq.TermTypeTry, Try: &Try{Body: &Query{Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "$x"}, SuffixList: []*Suffix{{Index: &Index{Name: "id"}}}}}, Op: gojq.OpEq, Right: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "id"}}}}}}, Catch: &Query{Func: "false"}}}}}, Then: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Func: "."}}}}, Else: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeQuery, Query: &Query{Left: &Query{Func: "child"}, Op: gojq.OpPipe, Right: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "lineage", Args: []*Query{{Func: "target"}, {Func: "child"}}}}}}}}, Op: gojq.OpAdd, Right: &Query{Term: &Term{Type: gojq.TermTypeArray, Array: &Array{Query: &Query{Func: "."}}}}}}}}}}}}}}, {Name: "lineage", Args: []string{"target"}, Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "lineage", Args: []*Query{{Func: "target"}, {Func: "children"}}}}}}, {Name: "lineage", Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "lineage", Args: []*Query{{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "isempty", Args: []*Query{{Func: "focus_child"}}}}}, {Func: "focus_child"}}}}}}, {Name: "scratchpad", Body: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "nodes"}, SuffixList: []*Suffix{{Iter: true}}}}, Op: gojq.OpPipe, Right: &Query{Left: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "select", Args: []*Query{{Left: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "name"}}}, Op: gojq.OpEq, Right: &Query{Term: &Term{Type: gojq.TermTypeString, Str: &String{Str: "__i3"}}}}}}}}, Op: gojq.OpPipe, Right: &Query{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "nodes"}, SuffixList: []*Suffix{{Index: &Index{Start: &Query{Term: &Term{Type: gojq.TermTypeNumber, Number: "0"}}}}}}}}}}, {Name: "tiles", Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "recurse", Args: []*Query{{Term: &Term{Type: gojq.TermTypeIndex, Index: &Index{Name: "nodes"}, SuffixList: []*Suffix{{Iter: true}}}}}}}}}, {Name: "leaves", Body: &Query{Term: &Term{Type: gojq.TermTypeFunc, Func: &Func{Name: "recurse", Args: []*Query{{Func: "children"}}}}}}}}
	default:
		return nil
	}
}