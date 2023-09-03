package Builders

import "builder/Products"

type Builder interface {
	SetLabtopType()
	GetLabtop() *Products.Labtop
}
