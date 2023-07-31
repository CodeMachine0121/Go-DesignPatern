package main

import (
	"fmt"
	file "prototype-practice/File"
	folder "prototype-practice/Folder"
	interfaces "prototype-practice/Interfaces"
)

func main() {

	// declare files
	file1 := file.File{Name: "File1"}
	file2 := file.File{Name: "File2"}
	file3 := file.File{Name: "File3"}

	// declare folders
	folder1 := folder.Folder{
		Name:     "Folder1",
		Children: []interfaces.InterfaceNode{&file1},
	}

	folder2 := folder.Folder{
		Name:     "Folder2",
		Children: []interfaces.InterfaceNode{&folder1, &file2, &file3},
	}

	fmt.Println("\n Print Folder2 Layer")
	folder2.Print(" ")

	cloneFolder2 := folder2.Clone()
	fmt.Println("\n Print Clone Folder2 Layer")
	cloneFolder2.Print(" ")
}
