package main

import (
	"builder-practice/Builder"
	"fmt"
)

func main() {

	builder := Builder.GetNewBuilder()
	builder.BuildLabtop("ROG")
	labtop := builder.GetLabtop()
	fmt.Println("New Labtop: ", labtop.GetName())
}
