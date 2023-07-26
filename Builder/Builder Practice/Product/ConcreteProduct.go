package Product

type ConcreteProduct struct {
	Name string
}

func (p *ConcreteProduct) GetName() string {
	return p.Name
}
