package main

import "fmt"

type Product interface {
	Operation()
}

type Product1 struct{}
type Product2 struct{}

func (p Product1) Operation() {
	fmt.Println("\tProduct1 Operation")
}

func (p Product2) Operation() {
	fmt.Println("\tProduct2 Operation")
}

type Creator interface {
	FactoryMethod() Product
}

type Creator1 struct{}
type Creator2 struct{}

func (c *Creator1) FactoryMethod() Product {
	return Product1{}
}

func (c *Creator2) FactoryMethod() Product {
	return Product2{}
}

func clientCode(c Creator) {
	p := c.FactoryMethod()
	p.Operation()
}

func main() {
	fmt.Println("Creator1:")
	clientCode(&Creator1{})
	fmt.Println("Creator2:")
	clientCode(&Creator2{})
}
