package Product

import "fmt"

type ConcreteProduct struct {
}

func (c *ConcreteProduct) PrintName() {
	fmt.Println("Get Concrete Product")
}
