package Flyweights

import "fmt"

type Flyweight interface {
	Operation()
}

type ConcretetFlyweight struct {
	intrinsicState string
}

func (fw ConcretetFlyweight) Init(intrinsicState string) {
	fw.intrinsicState = intrinsicState
}

func (fw ConcretetFlyweight) Operation(extringsicState string) string {
	fmt.Println("IntrinsicState: ", fw.intrinsicState)

	if extringsicState != "" {
		return extringsicState
	}

	return "Empty ExtrinsicState"
}

func NewConcreteFlyweight(state string) *ConcretetFlyweight {
	return &ConcretetFlyweight{
		intrinsicState: state,
	}
}
