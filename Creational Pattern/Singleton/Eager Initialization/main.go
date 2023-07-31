package main

import "fmt"

type Singleton struct {
	value int
}

var instance *Singleton

func init() {
	if instance == nil {
		fmt.Println("New instance of Singleton")
		instance = new(Singleton)
	}
}

func GetInstance() *Singleton {
	return instance
}

func main() {
	GetInstance()
}
