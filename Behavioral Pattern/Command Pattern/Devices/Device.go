// Package devices provides ...
package devices

import "fmt"

type Device interface {
	On()
	Off()
}

type Light struct {
	IsRunning bool
}

func (t *Light) On() {
	t.IsRunning = true
	fmt.Println("Toggle On")
}

func (t *Light) Off() {
	t.IsRunning = false
	fmt.Println("Toggle Off")
}
