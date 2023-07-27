package pools

import (
	"sync"
)

type Pool struct {
	sync.Mutex
	Inuse     []interface{}
	Available []interface{}
	New       func() interface{}
}

func NewPool(new func() interface{}) *Pool {
	return &Pool{New: new}
}

func (p *Pool) Acquire() interface{} {
	p.Lock()
	var object interface{}

	if len(p.Inuse) != 0 {
		object = p.Available[0]
		// TODO: remove one of Available
		p.Available = append(p.Available[:0], p.Available[1:]...)
		// TODO: add one in Inuse
		p.Inuse = append(p.Inuse, object)
	} else {
		object = p.New()
		p.Inuse = append(p.Inuse, object)
	}
	p.Unlock()
	return object
}
