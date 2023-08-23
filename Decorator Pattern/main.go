package main

import (
	"decorator/Components"
	decorators "decorator/Decorators"
)

func main() {

	concreteComponent := Components.ConcreteComponent{}

	decoratorA := decorators.DecoratorA{}
	decoratorB := decorators.DecoratorB{}

	// every decorator implements Component interface
	decoratorA.SetComponent(&concreteComponent)
	decoratorB.SetComponent(&decoratorA)
	decoratorB.Operation()

}
