package Builders

import "builder/Products"

type ConcreteBuilder struct {
	result Products.Product
}

func NewConcreteBuilder() ConcreteBuilder {
	return ConcreteBuilder{result: Products.Product{}}
}

func (b *ConcreteBuilder) Build() {
	b.result = Products.Product{}
}

func (b *ConcreteBuilder) GetResult() Products.Product {
	return Products.Product{Built: true}
}
