package Prototypes

type Prototype interface {
	Clone(string, string, string) Prototype
	GetType() string
}

type Labtop struct {
	Type    string
	Memory  string
	Storage string
}

// GetType implements Prototype.
func (l *Labtop) GetType() string {
	return l.Type
}

func (l *Labtop) Clone(Type, Memory, Storage string) Prototype {
	return &Labtop{Type, Memory, Storage}
}
