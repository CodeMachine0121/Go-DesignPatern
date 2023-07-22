package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

type Singleton struct {
	value int
}

var instance *Singleton

func GetInstance() *Singleton {
	mutex.Lock()
	defer mutex.Unlock()

	if instance == nil {
		instance = new(Singleton)
		fmt.Println(" New instance of Singleton")
	} else {
		fmt.Println(" Instance of Singleton Already Created")
	}
	return instance
}

func main() {
	for i := 0; i < 3; i++ {
		GetInstance()
	}
}
