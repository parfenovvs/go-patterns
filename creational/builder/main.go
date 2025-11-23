package main

import "fmt"

type Product struct {
	PartA string
	PartB string
}

type Director struct {
	Builder Builder
}

func (d Director) Construct() {
	d.Builder.BuildPartA()
	d.Builder.BuildPartB()
}

type Builder interface {
	BuildPartA()
	BuildPartB()
	GetResult() *Product
}

type ConcreteBuilder struct {
	p *Product
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{p: &Product{}}
}

func (b *ConcreteBuilder) BuildPartA() {
	b.p.PartA = "PartA"
}

func (b *ConcreteBuilder) BuildPartB() {
	b.p.PartB = "PartB"
}

func (b *ConcreteBuilder) GetResult() *Product {
	return b.p
}

func main() {
	var builder Builder = NewConcreteBuilder()
	director := Director{Builder: builder}
	director.Construct()
	product := builder.GetResult()

	fmt.Println("Product built with:", product.PartA, "and", product.PartB)
}
