package pools

func (p *Pool) Release(object interface{}) {
	p.Lock()
	p.Available = append(p.Available, object)
	for i, v := range p.Inuse {
		if v == object {
			// TODO: remove object from available list
			p.Inuse = append(p.Inuse[:i], p.Inuse[i+1:]...)
			break
		}
	}
	p.Unlock()
}
