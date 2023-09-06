package Labtop

import "sync"

type Labtop struct {
	Name string
}
type LabtopPool struct {
	sync.Mutex

	NotYetFashion []*Labtop
	OnFashion     []*Labtop
	OldFashion    []*Labtop
}

func InitLabtopPool(notYetFashion []*Labtop) *LabtopPool {
	return &LabtopPool{NotYetFashion: notYetFashion}
}

func (p *LabtopPool) AddToFashionList() {
	p.Lock()

	labtop := p.NotYetFashion[0]
	if len(p.OnFashion) != 0 {
		labotopNeedRetired := p.OnFashion[0]
		p.AddToOldFashionList(labotopNeedRetired)
		p.OnFashion = append(p.OnFashion[:0], p.OnFashion[1:]...)
	}
	p.OnFashion = append(p.OnFashion, labtop)
	p.NotYetFashion = append(p.NotYetFashion[:0], p.NotYetFashion[1:]...)

	p.Unlock()
}

func (p *LabtopPool) AddToOldFashionList(labotopNeedRetired *Labtop) {
	if len(p.OldFashion) != 0 {
		p.OldFashion = append(p.OldFashion[:0], p.OldFashion[1:]...)
	}
	p.OldFashion = append(p.OldFashion, labotopNeedRetired)
}
