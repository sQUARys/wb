package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type action string

const (
	A action = "A"
	B action = "B"
	C action = "C"
)

type Creator interface {
	CreateNewProduct(action action) Product // factory method
}

type Product interface {
	Use() string
}

type MemoryCreator struct{}

func NewCreator() Creator {
	return &MemoryCreator{}
}

func (m *MemoryCreator) CreateNewProduct(action action) Product {
	var product Product

	switch action {
	case A:
		product = &ProductA{action: string(action)}
	case B:
		product = &ProductB{action: string(action)}
	case C:
		product = &ProductC{action: string(action)}
	}
	return product
}

type ProductA struct {
	action string
}

func (p *ProductA) Use() string {
	return p.action
}

type ProductB struct {
	action string
}

func (p *ProductB) Use() string {
	return p.action
}

type ProductC struct {
	action string
}

func (p *ProductC) Use() string {
	return p.action
}

func main() {

	factory := NewCreator()
	products := []Product{
		factory.CreateNewProduct(A),
		factory.CreateNewProduct(B),
		factory.CreateNewProduct(C),
	}

	for _, product := range products {
		fmt.Println(product.Use())
	}

}
