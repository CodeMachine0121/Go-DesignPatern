package Products

type Product struct {
	Built bool
}

func NewProduct() *Product {
	return &Product{Built: false}
}
