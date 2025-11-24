package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type SomeSingleton struct {
	RandomNumber int
}

var instance *SomeSingleton
var once sync.Once

func GetInstance() *SomeSingleton {
	once.Do(func() {
		instance = &SomeSingleton{
			RandomNumber: rand.Int(),
		}
	})
	return instance
}

func main() {
	i1 := GetInstance()
	i2 := GetInstance()

	fmt.Println("i1: ", i1)
	fmt.Println("i2: ", i2)
}
