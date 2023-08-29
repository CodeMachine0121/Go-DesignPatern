package main

import (
	concretations "bridge/Concretations"
	refinds "bridge/Refinds"
)

func main() {
	concreteImplementor := concretations.NewConcreteImplementor()
	refindedAbstraction := refinds.NewRefindAbstraction(*concreteImplementor)
	refindedAbstraction.Execute("This is Bridge Mmode")
}
