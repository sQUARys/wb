package dev02

import (
	"fmt"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func unPack(str []rune) string {
	var memory rune
	isLetter := false
	var result string

	for _, word := range str {
		if unicode.IsDigit(word) && isLetter { // если нынешнее - число, предыдущее - строка
			number, _ := strconv.Atoi(string(word))
			for i := 0; i < number; i++ {
				result += string(memory)
			}
		} else if !unicode.IsDigit(word) && isLetter { // если предыдущее  - строка , нынешнее - строка
			result += string(memory)
		} else if unicode.IsDigit(word) && !isLetter { // если предыдущее -число, нынешнее - число
			fmt.Println("Не может стоять два числа подряд, пожалуйста проверьте ввод")
			break
		}

		memory = word
		if unicode.IsLetter(memory) {
			isLetter = true
		} else {
			isLetter = false
		}
	}
	if unicode.IsLetter(memory) {
		result += string(memory) // чтобы не потерять последний символ
	}
	return result
}

func main() {
	str := []rune("a4b3cdek5g9")
	out := unPack(str)

	fmt.Println(out)
}
