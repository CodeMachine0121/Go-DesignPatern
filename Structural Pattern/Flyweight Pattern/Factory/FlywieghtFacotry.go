package Factory

import "flyweight/Flyweights"

type FlyweightFactory struct {
	pool map[string]*Flyweights.ConcretetFlyweight
}

func NewFlyweightFactory() *FlyweightFactory {

	return &FlyweightFactory{pool: make(map[string]*Flyweights.ConcretetFlyweight)}
}

func (f *FlyweightFactory) GetFlyWeight(state string) *Flyweights.ConcretetFlyweight {
	flyweight, _ := f.pool[state]

	if f.pool[state] == nil {
		flyweight = Flyweights.NewConcreteFlyweight(state)
		f.pool[state] = flyweight
	}
	return flyweight
}
