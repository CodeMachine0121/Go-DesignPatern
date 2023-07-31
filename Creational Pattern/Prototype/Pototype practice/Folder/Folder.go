package folder

import (
	"fmt"
	interfaces "prototype-practice/Interfaces"
)

type Folder struct {
	Children []interfaces.InterfaceNode
	Name     string
}

func (f *Folder) Print(indetation string) {
	fmt.Println(indetation + "Folder Name: " + f.Name)

	for _, i := range f.Children {
		i.Print(indetation + indetation)
	}
}

func (f *Folder) Clone() interfaces.InterfaceNode {
	cloneFolder := &Folder{Name: "Clone-" + f.Name}

	var tempChildren []interfaces.InterfaceNode
	for _, i := range f.Children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}

	cloneFolder.Children = tempChildren
	return cloneFolder
}
