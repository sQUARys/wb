package main

import "fmt"

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	var value int64
	var i int

	fmt.Print("Введите значение: ")
	fmt.Scanf("%d\n", &value)
	fmt.Print("Введите какой бит заменить:")
	fmt.Scanf("%d\n", &i)

	fmt.Printf("Заменили в введенном числе %d бит и получили число: %d\n", value, value^(1<<i))
}
