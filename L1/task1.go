package main

import (
	"fmt"
)

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	YearOfBirth int
	Age         int
}

type location string

type Action struct {
	Human
	location
}

func (l location) loc() {
	fmt.Println("Сейчас вы находитесь здесь: ", l)
}

func (h Human) getDate() int {
	return (h.YearOfBirth + h.Age)
}

func main() {
	Action := Action{
		Human:    Human{YearOfBirth: 2000, Age: 21},
		location: "улица Ленина",
	}

	fmt.Println(Action.getDate())
	Action.loc()
}
