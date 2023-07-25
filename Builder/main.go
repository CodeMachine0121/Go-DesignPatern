package main

import (
	"builder/Builders"
	"builder/Directors"
	"fmt"
)

func main() {

	builder := Builders.NewConcreteBuilder()
	director := Directors.NewDirector(&builder)
	director.Construct()

	product := director.Builder.GetResult()
	fmt.Println(product)

}
