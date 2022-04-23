package main

import (
	"fmt"
)

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct { //структура с полями: возраст и год рождения
	YearOfBirth int
	Age         int
}

type location string // строка в которой хранится местоположение

type Action struct { // встраивание структуры Human
	Human
	location
}

func (l location) loc() {
	fmt.Println("Сейчас вы находитесь здесь: ", l)
}

func (h Human) getDate() int { // функция получения нынешнего года
	return (h.YearOfBirth + h.Age)
}

func main() {
	Action := Action{ // создание структуры
		Human:    Human{YearOfBirth: 2000, Age: 21},
		location: "улица Ленина",
	}

	fmt.Println(Action.getDate()) // в структуре Action вызывается встроенный метод структуры Human
	Action.loc()
}
