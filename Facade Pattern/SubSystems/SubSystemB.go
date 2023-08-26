package subsystems

import "fmt"

type SubSystemB struct{}

func NewSubSystemB(name string) *SubSystemB {
	return &SubSystemB{}
}

func (b *SubSystemB) MethodThree() {
	fmt.Println("SubSystemA - MethodThree")
}

func (b *SubSystemB) MethodFour() {
	fmt.Println("SubSystemA - MethodFour")
}
