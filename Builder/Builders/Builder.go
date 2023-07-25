package Builders

import "builder/Products"

type Builder interface {
	Build()
	GetResult() Products.Product
}
