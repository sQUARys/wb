package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var memDel string
	isDelimited := false

	configFile, err := ioutil.ReadFile("data1.txt")

	if err != nil {
		log.Fatal(err)
	}
	var configLines []string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	if len(text) == 0 {
		fmt.Println("Ошибка ввода.")
		return
	}

	arr := strings.Split(text, "-")

	for i := 1; i < len(arr); i++ {
		if strings.Contains(arr[i], "d") {
			fmt.Println("Delimiter...")
			dVal := strings.Index(arr[i], "'")
			delimiter := strings.Split(arr[i][dVal:], "'")
			delimiter = delimiter[1 : len(delimiter)-1]
			configLines = strings.Split(string(configFile), delimiter[0]) // разделяем каждое слово строки на элементы массива
			for j := range configLines {
				fmt.Println(configLines[j])
			}
			isDelimited = true
			memDel = delimiter[0]
		}
		if strings.Contains(arr[i], "f") {
			if !isDelimited {
				fmt.Println(string(configFile))
				configLines = strings.Split(string(configFile), "\t") // разделяем каждое слово строки на элементы массива
				fmt.Println(configLines[0])
			}

			fmt.Println("Fields...")
			var fields []string
			field, _ := strconv.Atoi(strings.Fields(arr[i])[1])

			if field >= len(configLines) {
				fmt.Println("Вы ввели поле, которое выходит за границы строк. Введите пожалуйста значение поменьше")
				return
			}
			for f := range configLines {
				fields = strings.Fields(configLines[f])
				fmt.Println(fields[field-1])
			}

		}
		if strings.Contains(arr[i], "s") {
			fmt.Println("Only-delimited...")

			configLines = strings.Split(string(configFile), "\n") // разделяем каждое слово строки на элементы массива
			if !isDelimited {
				memDel = "\t"
			}
			for s := range configLines {
				if strings.Contains(configLines[s], memDel) {
					fmt.Println(configLines[s])
				}
			}
		}
	}

}
