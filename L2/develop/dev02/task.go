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

//. Функция принимает строку. А в руны конвертирует внутри
//2. "В случае если была передана некорректная строка функция должна возвращать ошибку." у тебя функци яне возвращает ошибок
//3. Ты игнорируешь ошибки через _ - это нельзя делать. Это практически самое ужасное, что ты можешь делать в го. Никогда так не делай.
//4. for _, word := range str - ты итерируешься по строке. Но word не соответствует тому, что в нем лежит: там не слово.

func unPack(str string) string {
	var memory rune
	isLetter := false
	var result string

	strRune := []rune(str)

	for _, element := range strRune {
		if unicode.IsDigit(element) && isLetter { // если нынешнее - число, предыдущее - строка
			number, error := strconv.Atoi(string(element))
			if error != nil {
				fmt.Println("Error: ", error)
			}
			for i := 0; i < number; i++ {
				result += string(memory)
			}
		} else if !unicode.IsDigit(element) && isLetter { // если предыдущее - строка , нынешнее - строка
			result += string(memory)
		} else if unicode.IsDigit(element) && !isLetter { // если предыдущее -число, нынешнее - число
			result = ""
			break
		}

		memory = element
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
	str := "a4bc2d5e"
	out := unPack(str)

	fmt.Println(out)
}
