package Builder

import "builder-practice/Product"

type LabtopBuilder interface {
	Build()
	GetLabtop() Product.ConcreteProduct
}

type ConcreteBuilder struct {
	ResultProduct Product.ConcreteProduct
}

func GetNewBuilder() ConcreteBuilder {
	return ConcreteBuilder{ResultProduct: Product.ConcreteProduct{}}
}

func (p *ConcreteBuilder) BuildLabtop(labtopName string) {
	p.ResultProduct = Product.ConcreteProduct{Name: labtopName}
}

func (p *ConcreteBuilder) GetLabtop() Product.ConcreteProduct {
	return p.ResultProduct
}
