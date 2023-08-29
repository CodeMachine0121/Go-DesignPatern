package decorators

import (
	"decorator/Components"
	"fmt"
)

type DecoratorA struct {
	component Components.Component
}

func (da *DecoratorA) SetComponent(com Components.Component) {
	da.component = com
}

func (da *DecoratorA) Operation() {
	da.component.Operation()
	da.InpendentMethod()
}

func (da *DecoratorA) InpendentMethod() {
	fmt.Println("Inpendent Method of Decorator A")
}
