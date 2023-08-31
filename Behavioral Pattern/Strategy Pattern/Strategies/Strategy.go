package Strategies

import "fmt"

type Strategy interface {
	Execute()
}

type strategyA struct {
}

func (a *strategyA) Execute() {
	fmt.Println("Execute Strategy A")
}

//

type strategyB struct {
}

func (b *strategyB) Execute() {
	fmt.Println("Execute Strategy B")
}

// create Strategy Objects
func NewStrategyA() *strategyA {
	return &strategyA{}
}
func NewStrategyB() *strategyB {
	return &strategyB{}
}
