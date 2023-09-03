package Directors

import (
	"builder/Builders"
	"builder/Products"
)

type Director struct {
	buider Builders.Builder
}

func (d *Director) SetBuilder(buider Builders.Builder) {
	d.buider = buider
}

func NewDirector() *Director {
	return &Director{}
}

func (d *Director) BuilderLabtop() *Products.Labtop {
	return d.buider.GetLabtop()
}
