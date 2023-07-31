package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	value int
}

var once sync.Once
var instance *Singleton

func GetInstance() *Singleton {
	once.Do(func() {
		instance = new(Singleton)
		fmt.Println("New instance of Singleton")
	})
	return instance
}

func main() {
	GetInstance()
}
