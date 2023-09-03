package Directors

import "builder/Products"

type Builder interface {
	/* TODO: add methods */
	Build()
}

type ConcreteBuilder struct {
	result Products.Product
}

func NewConcreteBuilder() ConcreteBuilder {
	return ConcreteBuilder{}
}

func (builder *ConcreteBuilder) Build() {
	builder.result = *Products.NewProduct()
	builder.result.Built = true
}

func (builder *ConcreteBuilder) GetResult() Products.Product {
	return builder.result
}
