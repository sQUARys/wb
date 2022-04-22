package main

import (
	"fmt"
)

//Поменять местами два числа без создания временной переменной.

func changer(first int, second int) (int, int) {
	second = first + second
	first = second - first
	second = second - first
	return first, second
}

func main() {
	a := -10
	b := 5
	a, b = changer(a, b)

	fmt.Println(a, b)
}
