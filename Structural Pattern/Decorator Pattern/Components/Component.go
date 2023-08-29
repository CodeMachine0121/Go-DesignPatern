package Components

import "fmt"

type Component interface {
	Operation()
}

type ConcreteComponent struct {
}

func (c *ConcreteComponent) Operation() {
	fmt.Println("ConcreteComponent operating")
}
