package objects

type ObjectAdapter struct {
	Adaptee ObjectAdaptee
}

func (p *ObjectAdapter) Execute() {
	p.Adaptee.SpecifiExecute()
}
