package main

import (
	"implementation/Commands"
	"implementation/Invokers"
	"implementation/Receivers"
)

func main() {

	receiver := Receivers.NewReceiver()
	invoker := Invokers.NewInvoker()

	cmd1 := Commands.NewCommand1("Command 1", receiver)
	cmd2 := Commands.NewCommand2("Command 2", receiver)

	invoker.StoreCommand(cmd1)
	invoker.StoreCommand(cmd2)
	invoker.ExecuteCommands()
}
