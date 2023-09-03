package main

import (
	"builder/Builders"
	"builder/Directors"
	"fmt"
)

func main() {

	asusBuilder := Builders.NewAsusBuilder()
	rogBuilder := Builders.NewRogBuilder()

	director := Directors.NewDirector()

	director.SetBuilder(asusBuilder)
	asus := director.BuilderLabtop()
	fmt.Println("AsusBuilder: " + asus.LabtopType)

	director.SetBuilder(rogBuilder)
	rog := director.BuilderLabtop()
	fmt.Println("RogBuilder: " + rog.LabtopType)
}
