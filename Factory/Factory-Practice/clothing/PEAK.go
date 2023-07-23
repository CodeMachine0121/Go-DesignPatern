package clothing

type PEAK struct {
	Clothes
}

func NewPEAK() IClothes {
	return &PEAK{
		Clothes: Clothes{
			name: "New PEAK",
			size: 21,
		},
	}
}
