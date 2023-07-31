package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	value int
}

var instance *Singleton
var mutex sync.Mutex

func GetInstance() *Singleton {

	mutex.Lock()
	defer mutex.Unlock()
	if instance == nil {
		instance = new(Singleton)
		fmt.Println("new instance")
	}
	return instance
}

func main() {
	instance = GetInstance()
}
