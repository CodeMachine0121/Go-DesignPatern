package Commands

import "implementation/Receivers"

type Command interface {
	/* TODO: add methods */
	Execute()
}

type Command1 struct {
	name     string
	receiver *Receivers.Receiver
}

func NewCommand1(name string, receiver *Receivers.Receiver) *Command1 {
	return &Command1{name, receiver}
}
func (c *Command1) Execute() {
	c.receiver.Operation1(c.name)
}

type Command2 struct {
	name     string
	receiver *Receivers.Receiver
}

func NewCommand2(name string, receiver *Receivers.Receiver) *Command2 {
	return &Command2{name, receiver}
}
func (c *Command2) Execute() {
	c.receiver.Operation2(c.name, c.name, c.name)
}
