package clothing

type IClothes interface {
	SetName(name string)
	SetSize(size int)
	GetName() string
	GetSize() int
}
