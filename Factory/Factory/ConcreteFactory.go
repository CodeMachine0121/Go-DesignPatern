package Factory

type ConcreteFactory struct {
}

func (cf *ConcreteFactory) FactoryMethod(owner string) Product {
	p := &ConcreteProduct{}
	return p
}
