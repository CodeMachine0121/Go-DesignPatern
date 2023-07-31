package file

import (
	"fmt"
	interfaces "prototype-practice/Interfaces"
)

type File struct {
	Name string
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + "File name: " + f.Name)
}

func (f *File) Clone() interfaces.InterfaceNode {
	return &File{Name: "Clone file-" + f.Name}
}
