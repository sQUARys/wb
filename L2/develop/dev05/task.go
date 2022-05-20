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

func After(arr []string, substring string, inner []string) []string {
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

	start := make([]string, len(arr[:index+1]))
	end := make([]string, len(arr[index+1:]))
	copy(start, arr[:index+1])
	copy(end, arr[index+1:])

	for j := 0; j < len(inner); j++ {
		start = append(start, inner[j])
	}

	for i := 0; i < len(end); i++ {
		start = append(start, end[i])
	}
	return start
}

func Before(arr []string, substring string, inner []string) []string {
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

func Context(arr []string, substring string, leftSide []string, rightSide []string) []string {

	result := After(
		Before(arr, substring, leftSide),
		substring,
		rightSide,
	)

	return result
}

func CountLines(arr []byte) int {
	lines := strings.Split(string(arr), "\n")
	return len(lines)
}

func IgnoreCase(arr []string, substring string) int {
	count := 0

	for i := 0; i < len(arr); i++ {
		if strings.ToLower(arr[i]) == strings.ToLower(substring) {
			count++
		}
	}
	return count
}

func Invert(arr []string, substring string) (bool, []string) {
	isFind := false
	var index int

	for i := 0; i < len(arr); i++ {
		if arr[i] == substring {
			index = i
			isFind = true
		}
	}

	start := make([]string, len(arr[:index]))
	end := make([]string, len(arr[index+1:]))
	copy(start, arr[:index])
	copy(end, arr[index+1:])
	for j := 0; j < len(end); j++ {
		start = append(start, end[j])
	}

	return isFind, start

}

func Fixed(arr []string, subArr []string) {
	var count int

	for i := 0; i < len(arr); i++ {
		count = 0
		for j := 0; j < len(subArr); j++ {
			if arr[i] == subArr[j] {
				count++
				i++
			}
		}
		if count == len(subArr) {
			fmt.Println("Ваша строка была точно найдена, ее индекс: ", i)
			break
		}
	}
}

func GetLineNumber(arr []byte, substring string) int {
	lines := strings.Split(string(arr), "\n")
	var fieldLines []string
	isFind := false

	var result int

	for i := 0; i < len(lines); i++ {
		fieldLines = strings.Fields(lines[i])
		for j := 0; j < len(fieldLines); j++ {
			if fieldLines[j] == substring {
				result = i + 1
				isFind = true
				break
			}
		}
	}

	if !isFind {
		fmt.Println("Вашего слова нет в данном тексте")
		result = -1
	}
	return result
}

func insertSubstring() string {
	var sub string
	fmt.Print("Ваш файл был успешно считан. Введите пожалуйста слово, которое вы хотите найти: ")
	fmt.Scan(&sub)
	return sub
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

func main() {

	configFile, err := ioutil.ReadFile("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	configLines := strings.Fields(string(configFile)) // разделяем каждое слово строки на элементы массива

	var sub string

	in := []string{"строка1", "строка2"}

	subArr1 := []string{"слева1", "слева2"}
	subArr2 := []string{"справа1", "справа2"}

	forFixArr := "В большом тексте"

	commands := Command{}
	commands.setFlag()

	if commands.A == true { // "after" печатать +N строк после совпадения
		sub = insertSubstring()
		fmt.Println("Ваш текст был успешно преобразован:")
		configLines = After(configLines, sub, in)
		fmt.Println(configLines)
	} else if commands.B == true { // "before" печатать +N строк до совпадения
		sub = insertSubstring()
		fmt.Println("Ваш текст был успешно преобразован:")
		configLines = Before(configLines, sub, in)
		fmt.Println(configLines)
	} else if commands.C == true { // "context" (A+B) печатать ±N строк вокруг совпадения
		sub = insertSubstring()
		configLines = Context(configLines, sub, subArr1, subArr2)
		fmt.Println(configLines)
	} else if commands.c == true { // "count" (количество строк)
		fmt.Println("Количество строк: ", CountLines(configFile))
	} else if commands.i == true { //"ignore-case" (игнорировать регистр)
		sub = insertSubstring()
		fmt.Println("Количество совпадений без учета регистра: ", IgnoreCase(configLines, sub))
	} else if commands.v == true { //"invert" (вместо совпадения, исключать)
		sub = insertSubstring()
		isFindToInvert, _ := Invert(configLines, sub)
		if isFindToInvert {
			fmt.Println("Ваш текст после удаления введенной подстроки.")
			_, configLines = Invert(configLines, sub)
			fmt.Println(configLines)
		} else {
			fmt.Println("Не было найдено введенного слова для его удаления.")
		}
	} else if commands.F == true { //"fixed", точное совпадение со строкой, не паттерн
		Fixed(configLines, strings.Fields(forFixArr))
	} else if commands.n == true { //"line num", печатать номер строки
		sub = insertSubstring()
		fmt.Println("Ваше слово находится в строке номер: ", GetLineNumber(configFile, sub))
	} else {
		fmt.Println("Вы ввели несуществующую команду. Программа завершит работу...")
	}

}
