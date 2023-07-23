package clothing

type Clothes struct {
	name string
	size int
}

func (c *Clothes) SetName(name string) {
	c.name = name
}

func (c *Clothes) GetName() string {
	return c.name
}

func (c *Clothes) SetSize(size int) {
	c.size = size
}

func (c *Clothes) GetSize() int {
	return c.size
}
