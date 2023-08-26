package facades

import subsystems "facade/SubSystems"

type Facade struct {
	subSystemA subsystems.SubSystemA
	subSystemB subsystems.SubSystemB
}

func NewFacade() *Facade {
	return &Facade{
		subSystemA: subsystems.SubSystemA{},
		subSystemB: subsystems.SubSystemB{},
	}
}

func (c *Facade) MethodA() {
	c.subSystemB.MethodThree()
	c.subSystemA.MethodOne()
	c.subSystemB.MethodFour()
}

func (c *Facade) MethodB() {
	c.subSystemB.MethodFour()
	c.subSystemA.MethodTwo()
}
