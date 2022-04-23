package main

import "fmt"

//Разработать программу, которая переворачивает подаваемую на ход строку (
//например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {
	s := "главрыба"                                       // строка-значение
	runes := []rune(s)                                    // создаем руну, а не строку, так как символы могут быть unicode
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { //i , j - два крайних значения
		runes[i], runes[j] = runes[j], runes[i] // меняем местами крайние значения и меняем крайние значения
	}
	fmt.Println(string(runes)) //выводим руну
}
