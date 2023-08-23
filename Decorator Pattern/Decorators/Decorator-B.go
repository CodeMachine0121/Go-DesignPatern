package decorators

import (
	"decorator/Components"
	"fmt"
)

type DecoratorB struct {
	component Components.Component
}

func (db *DecoratorB) SetComponent(component Components.Component) {
	db.component = component
}

func (db *DecoratorB) Operation() {
	db.component.Operation()
	fmt.Println(db.String())
}

func (db *DecoratorB) String() string {
	return "Independent method of DecoratorB"
}
