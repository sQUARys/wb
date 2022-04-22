package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке уникальные
//(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

func main() {
	s := "aabcd"
	s = strings.ToLower(s)
	isUnique := true

	for i := 0; i < len(s) && isUnique; i++ {
		if strings.Count(s, string(s[i])) > 1 {
			isUnique = false
		}
	}
	fmt.Println(isUnique)
}
