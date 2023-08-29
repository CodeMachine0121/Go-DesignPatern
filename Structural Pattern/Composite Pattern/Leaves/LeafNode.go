package Leaves

import (
	components "composit-pattern/Components"
	"fmt"
)

type Leaf struct {
	value int
}

func NewLeaf(value int) *Leaf {
	return &Leaf{value: value}
}

func (l *Leaf) Excute() {
	fmt.Println(l.value)
}

func (l *Leaf) Add(c components.Component) {
}
