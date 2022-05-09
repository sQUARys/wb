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
	configFile, err := ioutil.ReadFile("data.txt")

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
			dVal := strings.Index(arr[i], "'")
			//fmt.Println(string(arr[i][val]), "he")
			delimiter := strings.Split(arr[i][dVal:], "'")
			fmt.Println(delimiter)
			delimiter = delimiter[1 : len(delimiter)-1]
			fmt.Println(delimiter[0])
			//как разделить по табу
			configLines = strings.Split(string(configFile), delimiter[0]) // разделяем каждое слово строки на элементы массива
			for j := range configLines {
				fmt.Println(configLines[j])
			}
		}
		if strings.Contains(arr[i], "f") {
			var fields []string
			fmt.Println(arr[i])
			field, _ := strconv.Atoi(strings.Fields(arr[i])[1])
			for f := range configLines {
				fields = strings.Fields(configLines[f])
				fmt.Println(fields[field-1])
			}
		}
	}

}
