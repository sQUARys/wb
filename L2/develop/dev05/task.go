package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)
Поддержать флаги:
-A - "after" печатать +N строк после совпадения Done
-B - "before" печатать +N строк до совпадения Done
-C - "context" (A+B) печатать ±N строк вокруг совпадения Done
-c - "count" (количество строк) Done
-i - "ignore-case" (игнорировать регистр) Done
-v - "invert" (вместо совпадения, исключать) Done
-F - "fixed", точное совпадение со строкой, не паттерн Done
-n - "line num", печатать номер строки Done

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {

	configFile, err := ioutil.ReadFile("data.txt") // считываем данные из файла

	if err != nil {
		log.Fatal(err)
	}

	configLines := strings.Fields(string(configFile)) // разделяем каждое слово строки на элементы массива

	var sub string

	// устанавливаем какие слова хотим вставить
	in := []string{"строка1", "строка2"}
	subArr1 := []string{"слева1", "слева2"}
	subArr2 := []string{"справа1", "справа2"}

	forFixArr := "В большом тексте"

	commands := Command{}
	commands.setFlag()

	switch {
	case commands.A: //"after" печатать +N строк после совпадения Done
		sub = insertSubstring() // считываем ввод пользователя
		configLines = after(configLines, sub, in)
		fmt.Println(configLines)
	case commands.B: // "before" печатать +N строк до совпадения
		sub = insertSubstring() // считываем ввод пользователя
		configLines = before(configLines, sub, in)
		fmt.Println(configLines)
	case commands.C: // "context" (A+B) печатать ±N строк вокруг совпадения
		sub = insertSubstring() // считываем ввод пользователя
		configLines = context(configLines, sub, subArr1, subArr2)
		fmt.Println(configLines)
	case commands.c: // "count" (количество строк)
		fmt.Println(countLines(configFile))
	case commands.i: //"ignore-case" (игнорировать регистр)
		sub = insertSubstring() // считываем ввод пользователя
		fmt.Println(ignoreCase(configLines, sub))
	case commands.v: //"invert" (вместо совпадения, исключать)
		sub = insertSubstring() // считываем ввод пользователя
		isFindToInvert, _ := invert(configLines, sub)
		if isFindToInvert {
			_, configLines = invert(configLines, sub)
			fmt.Println(configLines)
		}
	case commands.F: //"fixed", точное совпадение со строкой, не паттерн
		fixed(configLines, strings.Fields(forFixArr))
	case commands.n: //"line num", печатать номер строки
		sub = insertSubstring() // считываем ввод пользователя
		number, err := getLineNumber(configFile, sub)
		if err {
			break
		}
		fmt.Println(number)
	default:
		fmt.Println("Вы ввели несуществующую команду. Программа завершит работу...")
	}
}

type Command struct {
	A bool
	B bool
	C bool
	c bool
	i bool
	v bool
	F bool
	n bool
}

func (c *Command) setFlag() {
	flag.BoolVar(&c.A, "A", false, "after")
	flag.BoolVar(&c.B, "B", false, "before")
	flag.BoolVar(&c.C, "C", false, "context")
	flag.BoolVar(&c.c, "c", false, "countlines")
	flag.BoolVar(&c.i, "i", false, "ignore-case")
	flag.BoolVar(&c.v, "v", false, "invert")
	flag.BoolVar(&c.F, "F", false, "fixed")
	flag.BoolVar(&c.n, "n", false, "get line numb")
	flag.Parse()
}

func after(arr []string, substring string, inner []string) []string {
	index := -1 // переменная для хранения индекса

	for i := 0; i < len(arr); i++ {
		if arr[i] == substring { // если найдено слово, которое искал пользователь
			index = i
			break
		}
	}
	if index == -1 { // если слово не найдено
		return []string{""}
	}

	start := make([]string, len(arr[:index+1])) // делим массив на два: до слова и после
	end := make([]string, len(arr[index+1:]))
	copy(start, arr[:index+1]) // избавляемся от зависимости
	copy(end, arr[index+1:])

	for j := 0; j < len(inner); j++ {
		start = append(start, inner[j]) // заполняем массив со вставками вместе "после"
	}

	for i := 0; i < len(end); i++ {
		start = append(start, end[i]) // дополняем массив оставшимися элементами
	}
	return start
}

func before(arr []string, substring string, inner []string) []string {
	index := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == substring {
			index = i
			break
		}
	}
	if index == -1 {
		return []string{""}
	}

	start := make([]string, len(arr[:index-len(inner)+2]))
	end := make([]string, len(arr[index:]))
	copy(start, arr[:index-len(inner)+2])
	copy(end, arr[index:])

	for j := 0; j < len(inner); j++ {
		start = append(start, inner[j])
	}

	for i := 0; i < len(end); i++ {
		start = append(start, end[i])
	}
	return start
}

func context(arr []string, substring string, leftSide []string, rightSide []string) []string {
	result := after(
		before(arr, substring, leftSide),
		substring,
		rightSide,
	) // делаем двойную вставку до и после с уже написанными функциями

	return result
}

func countLines(arr []byte) int {
	lines := strings.Split(string(arr), "\n") // делим массив по \n
	return len(lines)                         // считаем длину строк
}

func ignoreCase(arr []string, substring string) int {
	count := 0

	for i := 0; i < len(arr); i++ {
		if strings.ToLower(arr[i]) == strings.ToLower(substring) { // сравниваем слова приведенные к нижнему регистру
			count++
		}
	}
	return count
}

func invert(arr []string, substring string) (bool, []string) {
	isFind := false
	var index int

	for i := 0; i < len(arr); i++ {
		if arr[i] == substring { // нашли искомое слово
			index = i
			isFind = true
		}
	}

	start := make([]string, len(arr[:index])) // делим массив на два до и после слова не включая его
	end := make([]string, len(arr[index+1:]))
	copy(start, arr[:index]) // избавляемся от зависимостей
	copy(end, arr[index+1:])
	for j := 0; j < len(end); j++ {
		start = append(start, end[j]) // соединяем два массива без найденного слова, для его удаленияя
	}

	return isFind, start

}

func fixed(arr []string, subArr []string) {
	var count int

	for i := 0; i < len(arr); i++ {
		count = 0
		for j := 0; j < len(subArr); j++ {
			if arr[i] == subArr[j] { // поиск всей строки полностью
				count++ // считаем количество слов
				i++
			}
		}
		if count == len(subArr) { // если количество слов совпадает с их количеством в пользовательском вводе
			fmt.Println("Ваша строка была точно найдена, ее индекс: ", i)
			break
		}
	}
}

func getLineNumber(arr []byte, substring string) (int, bool) {
	lines := strings.Split(string(arr), "\n") // делим массив по \n
	var fieldLines []string
	isFind := false

	var result int
	err := false

	for i := 0; i < len(lines); i++ {
		fieldLines = strings.Fields(lines[i]) // не учитываем пропуски
		for j := 0; j < len(fieldLines); j++ {
			if fieldLines[j] == substring { // если найдена строка
				result = i + 1 //найденный индекс записываем в переменную
				isFind = true  // нашли слово
				break
			}
		}
	}

	if !isFind {
		fmt.Println("Вашего слова нет в данном тексте")
		err = true
	}
	return result, err
}

func insertSubstring() string {
	var sub string
	fmt.Print("Ваш файл был успешно считан. Введите пожалуйста слово, которое вы хотите найти: ")
	fmt.Scan(&sub)
	return sub
}
