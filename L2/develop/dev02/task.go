package main

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

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint. DONE
*/

func unPack(str string) string { // функция развертки данных
	var memory rune
	isLetter := false
	var result string

	strRune := []rune(str)

	for _, element := range strRune { // проходим по всей руне
		if unicode.IsDigit(element) && isLetter { // если нынешнее - число, предыдущее - строка
			number, error := strconv.Atoi(string(element)) // переводим строку в число
			if error != nil {
				fmt.Println("Error: ", error)
			}
			for i := 0; i < number; i++ {
				result += string(memory) // печатаем столько раз сколько пользователь ввел
			}
		} else if !unicode.IsDigit(element) && isLetter { // если предыдущее - строка , нынешнее - строка
			result += string(memory) // просто складываем два элемента в строку
		} else if unicode.IsDigit(element) && !isLetter { // если предыдущее -число, нынешнее - число
			result = "" // не может стоять два числа подряд
			break
		}
		memory = element
		if unicode.IsLetter(memory) { // если нынешнее было числом
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
	str := "a4bc2d5e"
	out := unPack(str) // распаковка

	fmt.Println(out)
}
