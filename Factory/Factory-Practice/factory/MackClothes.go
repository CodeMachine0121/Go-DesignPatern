package factory

import (
	"factory-practice/clothing"
	"fmt"
)

func MackeClothes(clothingType string) (clothing.IClothes, error) {

	if clothingType == "ANTA" {
		return clothing.NewANTA(), nil
	}

	if clothingType == "PEAK" {
		return clothing.NewPEAK(), nil
	}

	return nil, fmt.Errorf("Invalid clothing type: %s", clothingType)
}
