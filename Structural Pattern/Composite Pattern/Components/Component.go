package Components

type Component interface {
	Excute()
	Add(Component)
}
