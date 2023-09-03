package Invokers

import "implementation/Commands"

type Invoker struct {
	cmds []Commands.Command
}

func (i *Invoker) StoreCommand(cmd Commands.Command) {
	i.cmds = append(i.cmds, cmd)
}

func (i *Invoker) ExecuteCommands() {
	for _, cmd := range i.cmds {
		cmd.Execute()
	}
}

func NewInvoker() *Invoker {
	return &Invoker{}
}
