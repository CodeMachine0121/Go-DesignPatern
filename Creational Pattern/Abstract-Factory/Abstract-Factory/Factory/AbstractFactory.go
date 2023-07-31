package Factory

import "abFactory/Product"

type AbstractFactory interface {
	CreateProduct() Product.AbstractProduct
}
