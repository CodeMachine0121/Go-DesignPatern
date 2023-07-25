package Directors

import "builder/Builders"

type Director struct {
	Builder Builders.Builder
}

func NewDirector(builder Builders.Builder) Director {
	return Director{Builder: builder}
}

func (d *Director) Construct() {
	d.Builder.Build()
}
