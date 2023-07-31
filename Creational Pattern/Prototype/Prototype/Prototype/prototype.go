package prototype

type Prototype interface {
	GetName() string
	Clone() Prototype
}

type ConcretePrototype struct {
	Name string
}

func (p *ConcretePrototype) GetName() string { return p.Name }

func (p *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{p.Name}
}
