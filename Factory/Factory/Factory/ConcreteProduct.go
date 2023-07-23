package Factory

import "fmt"

type ConcreteProduct struct {
}

func (p *ConcreteProduct) Use() {
	fmt.Println("Use Concrete Product")
}
