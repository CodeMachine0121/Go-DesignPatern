package main

import (
	"fmt"
	prototype "prototype/Prototype"
)

func main() {

	proto := prototype.ConcretePrototype{Name: "Concrete Prototype"}

	clone := proto.Clone("Concrete Prototype Clone")

	fmt.Println("proto: ", proto.GetName())
	fmt.Println("clone: ", clone.GetName())

}
