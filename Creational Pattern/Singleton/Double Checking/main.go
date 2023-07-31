package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	value int
}

var mutex sync.Mutex
var instance *Singleton

func GetInstance() *Singleton {
	if instance == nil {
		mutex.Lock()
		if instance == nil {
			instance = new(Singleton)
			fmt.Println("New Singleton")
		}
		mutex.Unlock()
	}
	return instance
}

func main() {
	GetInstance()
}
