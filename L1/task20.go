package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func main() {
	s := " dog sun snow help need" // строка

	words := strings.Fields(s) // разделить по словам в независимости от пробелов между ними

	var memory string // переменная для запоминания промежуточной строки

	j := len(words) - 1 // крайнее значение массива

	for i := 0; i < len(words) && i < j; i++ { // перебираем весь массив
		memory = words[i]   // запоминаем значение words[i]
		words[i] = words[j] // меняем значение левой границы со значением правой границы
		words[j] = memory   // записываем в значение правой границы записанное ранее значение
		j--                 // уменьшаем границы
	}

	for i := 0; i < len(words); i++ {
		fmt.Printf("%s ", words[i]) // выводим массив
	}

}
