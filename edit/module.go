package edit

import (
	"github.com/elves/elvish/eval"
	"github.com/elves/elvish/util"
)

// Interface between the editor and elvish script. Implements the le: module.

// makeModule builds a module from an Editor.
func makeModule(ed *Editor) eval.Namespace {
	ns := eval.Namespace{}
	// Populate builtins.
	for _, b := range builtins {
		ns[eval.FnPrefix+b.name] = eval.NewPtrVariable(&BuiltinAsFnValue{b, ed})
	}
	// Populate binding tables in the variable $binding.
	// TODO Make binding specific to the Editor.
	binding := &eval.Struct{
		[]string{"insert", "command", "completion", "navigation", "history"},
		[]eval.Variable{
			eval.NewRoVariable(BindingTable{keyBindings[modeInsert]}),
			eval.NewRoVariable(BindingTable{keyBindings[modeCommand]}),
			eval.NewRoVariable(BindingTable{keyBindings[modeCompletion]}),
			eval.NewRoVariable(BindingTable{keyBindings[modeNavigation]}),
			eval.NewRoVariable(BindingTable{keyBindings[modeHistory]}),
		},
	}

	ns["binding"] = eval.NewRoVariable(binding)
	ns["completer"] = eval.NewRoVariable(CompleterTable(argCompleter))

	ns["prompt"] = PromptVariable{&ed.ps1}
	ns["rprompt"] = PromptVariable{&ed.rps1}

	return ns
}

func throw(e error) {
	util.Throw(e)
}
