package main

import (
	"fmt"
	"maps"
	"slices"
)

type Prototype interface {
	Clone() Prototype
}

type ConcretePrototype1 struct {
	Data1 map[string]string
}

type ConcretePrototype2 struct {
	Data2 []int
}

func (c *ConcretePrototype1) Clone() Prototype {
	return &ConcretePrototype1{
		Data1: maps.Clone(c.Data1),
	}
}

func (c *ConcretePrototype2) Clone() Prototype {
	return &ConcretePrototype2{
		Data2: slices.Clone(c.Data2),
	}
}

func main() {
	prototype1 := &ConcretePrototype1{Data1: map[string]string{"key": "value"}}
	clone1 := prototype1.Clone().(*ConcretePrototype1)

	prototype2 := &ConcretePrototype2{Data2: []int{1, 2, 3}}
	clone2 := prototype2.Clone().(*ConcretePrototype2)

	fmt.Println("Original Prototype1 Data:", prototype1.Data1["key"])
	fmt.Println("Cloned Prototype1 Data:", clone1.Data1["key"])

	fmt.Println("Original Prototype2 Data:", prototype2.Data2[0])
	fmt.Println("Cloned Prototype2 Data:", clone2.Data2[0])

	prototype1.Data1["key"] = "new value"
	fmt.Println("Original Prototype1 Data:", prototype1.Data1["key"])
	fmt.Println("Cloned Prototype1 Data:", clone1.Data1["key"])
}
