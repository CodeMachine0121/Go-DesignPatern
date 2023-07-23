package clothing

type ANTA struct {
	Clothes
}

func NewANTA() IClothes {
	return &ANTA{
		Clothes: Clothes{
			name: "New ANTA",
			size: 12,
		},
	}
}
