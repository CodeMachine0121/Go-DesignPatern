package main

import (
	facades "facade/Facades"
	subsystems "facade/SubSystems"
)

func main() {
	fa := facades.NewFacade()
	fa.MethodA()
	fa.MethodB()

	sub := subsystems.NewSubSystemA()
	sub.MethodOne()
	sub.MethodTwo()
}
