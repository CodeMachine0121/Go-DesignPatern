package main

import "abFactory/Factory"

func main() {
	factory := Factory.NewConcreteFactory()
	product := factory.CreateProduct()
	product.PrintName()
}
