package main

import (
	"builder/Directors"
	"fmt"
)

func main() {

	builder := Directors.ConcreteBuilder{}
	director := Directors.NewDirector(&builder)
	director.Construct()

	product := builder.GetResult()
	fmt.Println(product)

}
