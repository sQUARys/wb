package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func main() {
	s := " dog sun snow help need"

	words := strings.Fields(s) // разделить по словам в независимости от пробелов между ними
	var memory string
	j := len(words) - 1

	for i := 0; i < len(words) && i < j; i++ {
		memory = words[i]
		words[i] = words[j]
		words[j] = memory
		j--
	}

	for i := 0; i < len(words); i++ {
		fmt.Printf("%s ", words[i])
	}

}
