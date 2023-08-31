package main

import (
	"strategy/Contexts"
	"strategy/Strategies"
)

func main() {
	strategyB := Strategies.NewStrategyB()
	context := Contexts.NewContext()
	context.SetStrategy(strategyB)

	strategyA := Strategies.NewStrategyA()
	context.SetStrategy(strategyA)

	context.Execute()
}
