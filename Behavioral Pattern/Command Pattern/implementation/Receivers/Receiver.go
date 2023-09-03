package Receivers

import "fmt"

type Receiver struct{}

func NewReceiver() *Receiver {
	return &Receiver{}
}

func (r *Receiver) Operation1(a string) {

	fmt.Println("Receiver.Operation1", a)
}

func (r *Receiver) Operation2(a, b, c string) {

	fmt.Println("Receiver.Operation2", a, b, c)
}
