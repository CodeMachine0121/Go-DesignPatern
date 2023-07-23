package main

import "factory/Factory"

func main() {
	fact := Factory.ConcreteFactory{}
	product := fact.FactoryMethod("James")
	product.Use()
}
