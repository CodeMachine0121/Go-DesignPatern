package Builders

import "builder/Products"

type RogBuilder struct {
	LabtopType string
}

func (rog *RogBuilder) SetLabtopType() {
	rog.LabtopType = "ROG"
}

func (rog *RogBuilder) GetLabtop() *Products.Labtop {
	rog.SetLabtopType()
	return &Products.Labtop{LabtopType: rog.LabtopType}
}

func NewRogBuilder() *RogBuilder {
	return &RogBuilder{}
}
