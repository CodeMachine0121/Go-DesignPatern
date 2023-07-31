package Factory

import "abFactory/Product"

type ConcreteFactory struct {
}

func NewConcreteFactory() ConcreteFactory {
	return ConcreteFactory{}
}

func (s *ConcreteFactory) CreateProduct() Product.ConcreteProduct {
	return Product.ConcreteProduct{}
}
