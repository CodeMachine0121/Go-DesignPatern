package refinds

import concretations "bridge/Concretations"

type RefindAbstraction struct {
	method concretations.ConcreteImplementor
}

func (c *RefindAbstraction) Execute(str string) {
	c.method.Implementation(str)
}

func NewRefindAbstraction(implementor concretations.ConcreteImplementor) *RefindAbstraction {
	return &RefindAbstraction{
		method: implementor,
	}
}
