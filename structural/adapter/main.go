package main

import "fmt"

type Target interface {
	Operation() string
}

type Adaptee struct{}

func (a *Adaptee) SpecificOperation() string {
	return "Adaptee: Specific operation"
}

type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Operation() string {
	return "Adapter: (TRANSLATED) " + a.adaptee.SpecificOperation()
}

func clientCode(t Target) {
	result := t.Operation()
	fmt.Println(result)
}

func main() {
	adaptee := &Adaptee{}
	target := &Adapter{adaptee: adaptee}
	clientCode(target)
}
