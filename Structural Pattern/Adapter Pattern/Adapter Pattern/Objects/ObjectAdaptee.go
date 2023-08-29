package objects

import "fmt"

type ObjectAdaptee struct {
}

func (b *ObjectAdaptee) SpecifiExecute() {
	fmt.Println("Executing: SpecificExecute")
}
