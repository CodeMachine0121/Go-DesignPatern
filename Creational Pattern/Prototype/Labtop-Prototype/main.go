package main

import (
	"fmt"
	"labtop/Prototypes"
)

func main() {
	labtopPrototype := &Prototypes.Labtop{}

	labotp := labtopPrototype.Clone("Rog", "15GB", "1TB")
	fmt.Println(labotp)
	Dell := labtopPrototype.Clone("Dell", "8GB", "500GB")
	fmt.Println(Dell)
	fmt.Println(labotp)
}
