package subsystems

import "fmt"

type SubSystemA struct{}

func NewSubSystemA() *SubSystemA {
	return &SubSystemA{}
}

func (s *SubSystemA) MethodOne() {
	fmt.Println("SubSystemB: MethodOne")
}

func (s *SubSystemA) MethodTwo() {
	fmt.Println("SubSystemB: MethodTwo")
}
