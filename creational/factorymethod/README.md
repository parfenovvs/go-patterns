# Factory Method Pattern

## Overview

The Factory Method pattern defines an interface for creating an object, but lets subclasses decide which class to instantiate. It lets a class defer instantiation to subclasses.

## Structure

```mermaid
classDiagram
    class Product {
        <<interface>>
        +Operation()
    }

    class Product1 {
        +Operation()
    }

    class Product2 {
        +Operation()
    }

    class Creator {
        <<interface>>
        +FactoryMethod() Product
    }

    class Creator1 {
        +FactoryMethod() Product
    }

    class Creator2 {
        +FactoryMethod() Product
    }

    Product <|.. Product1
    Product <|.. Product2
    Creator <|.. Creator1
    Creator <|.. Creator2
    Creator1 ..> Product1 : creates
    Creator2 ..> Product2 : creates
```

## Usage

```go
func main() {
    fmt.Println("Creator1:")
    clientCode(&Creator1{})
    fmt.Println("Creator2:")
    clientCode(&Creator2{})
}
```
