package concretations

import "fmt"

type ConcreteImplementor struct{}

func (c *ConcreteImplementor) Implementation(str string) {
	fmt.Println("print message: ", str)
}

func NewConcreteImplementor() *ConcreteImplementor {
	return &ConcreteImplementor{}
}
