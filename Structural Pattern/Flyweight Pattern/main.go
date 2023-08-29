package main

import (
	"flyweight/Factory"
	"fmt"
)

func main() {
	factory := Factory.NewFlyweightFactory()
	flyweight1 := factory.GetFlyWeight("fly-1")
	flyweight2 := factory.GetFlyWeight("fly-2")

	fmt.Println(flyweight1.Operation("ready"))
	fmt.Println(flyweight2.Operation("ok"))
}
