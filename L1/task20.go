package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

//Это конечно работает, но неудобно. Нужна функция, которая на вход принимает строку и на выход отдает строку.

func swap(str []string) []string {
	j := len(str) - 1 // крайнее значение массива

	for i := 0; i < len(str) && i < j; i++ { // перебираем весь массив
		memory := str[i] // запоминаем значение words[i]
		str[i] = str[j]  // меняем значение левой границы со значением правой границы
		str[j] = memory  // записываем в значение правой границы записанное ранее значение
		j--              // уменьшаем границы
	}
	return str
}

func main() {
	s := " dog sun snow help need" // строка

	words := strings.Fields(s) // разделить по словам в независимости от пробелов между ними

	words = swap(words) //меняем местами слова

	for i := 0; i < len(words); i++ {
		fmt.Printf("%s ", words[i]) // выводим массив
	}

}
