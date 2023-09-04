package prototype

type Prototype interface {
	GetName() string
	Clone(name string) Prototype
}

type ConcretePrototype struct {
	Name string
}

func (p *ConcretePrototype) GetName() string { return p.Name }

func (p *ConcretePrototype) Clone(name string) Prototype {
	return &ConcretePrototype{p.Name + " " + name}
}
