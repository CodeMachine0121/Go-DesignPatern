package composites

import (
	components "composit-pattern/Components"
)

type Composite struct {
	children []components.Component
}

func NewComposite() *Composite {
	// TODO: return  a zero-length composite
	return &Composite{make([]components.Component, 0)}
}

func (c *Composite) Add(component components.Component) {
	c.children = append(c.children, component)
}

func (c *Composite) Excute() {
	for i := 0; i < len(c.children); i++ {
		c.children[i].Excute()
	}
}
