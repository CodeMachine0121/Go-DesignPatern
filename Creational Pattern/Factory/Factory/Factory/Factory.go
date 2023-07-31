package Factory

type Factory interface {
	FactoryMethod(owner string) Product
}
