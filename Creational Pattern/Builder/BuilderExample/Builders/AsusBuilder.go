package Builders

import "builder/Products"

type AsusBuilder struct {
	LabtopType string
}

func (a *AsusBuilder) SetLabtopType() {
	a.LabtopType = "ASUS"
}

func (a *AsusBuilder) GetLabtop() *Products.Labtop {
	a.SetLabtopType()
	return &Products.Labtop{LabtopType: a.LabtopType}
}

func NewAsusBuilder() *AsusBuilder {
	return &AsusBuilder{}
}
