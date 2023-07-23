package main

import (
	"factory-practice/clothing"
	"factory-practice/factory"
	"fmt"
)

func main() {
	ANTA, _ := factory.MackeClothes("ANTA")
	PEAK, _ := factory.MackeClothes("PEAK")

	PrintInformation(ANTA)
	PrintInformation(PEAK)
}

func PrintInformation(Cloth clothing.IClothes) {
	fmt.Println("Clothes: ", Cloth.GetName())
	fmt.Println("Size: ", Cloth.GetSize())
	fmt.Println()
}
