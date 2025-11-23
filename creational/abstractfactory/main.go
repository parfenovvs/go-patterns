package main

import "fmt"

type AbstractProductA interface {
	Name() string
}
type ConcreteProductA struct {
	name string
}

func (p ConcreteProductA) Name() string {
	return p.name
}

type AbstractProductB interface {
	Title() string
}
type ConcreteProductB struct {
	title string
}

func (p ConcreteProductB) Title() string {
	return p.title
}

type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}
type ConcreteFactory1 struct{}
type ConcreteFactory2 struct{}

func (f ConcreteFactory1) CreateProductA() AbstractProductA {
	return ConcreteProductA{
		name: "ConcreteProductA1",
	}
}

func (f ConcreteFactory1) CreateProductB() AbstractProductB {
	return ConcreteProductB{
		title: "ConcreteProductB1",
	}
}

func (f ConcreteFactory2) CreateProductA() AbstractProductA {
	return ConcreteProductA{
		name: "ConcreteProductA2",
	}
}

func (f ConcreteFactory2) CreateProductB() AbstractProductB {
	return ConcreteProductB{
		title: "ConcreteProductB2",
	}
}

func clientCode(factory AbstractFactory) {
	productA := factory.CreateProductA()
	productB := factory.CreateProductB()
	fmt.Println("\tProduct A:", productA.Name())
	fmt.Println("\tProduct B:", productB.Title())
}

func main() {
	fmt.Println("Factory 1:")
	clientCode(ConcreteFactory1{})

	fmt.Println("Factory 2:")
	clientCode(ConcreteFactory2{})
}
