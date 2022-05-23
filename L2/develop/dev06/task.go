package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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

func (c *Command) Delimit(configFile []byte) []string {
	fmt.Println("Delimiter...")
	fmt.Println(string(configFile), "'", c.d, "'")
	configLines := strings.Split(string(configFile), c.d) // разделяем каждое слово строки на элементы массива
	c.isDelimited = true
	for j := range configLines {
		fmt.Println(configLines[j])
	}
	return configLines
}

//func Fields(i int, arr []string, configLines []string) bool {
//	fmt.Println("Fields...")
//	var fields []string
//	field, _ := strconv.Atoi(strings.Fields(arr[i])[1])
//
//	if field >= len(configLines) {
//		fmt.Println("Вы ввели поле, которое выходит за границы строк. Введите пожалуйста значение поменьше")
//		return
//	}
//	for f := range configLines {
//		fields = strings.Fields(configLines[f])
//		fmt.Println(fields[field-1])
//	}
//}

type Command struct {
	d           string
	f           bool
	s           bool
	isDelimited bool
}

func (c *Command) setFlags() {
	flag.StringVar(&c.d, "d", "\t", "delimiter")
	flag.BoolVar(&c.f, "f", false, "fields")
	flag.BoolVar(&c.s, "s", false, "separated")
	flag.Parse()
}

func main() {
	commands := Command{}
	commands.setFlags()
	commands.isDelimited = false

	configFile, err := ioutil.ReadFile("data1.txt")

	if err != nil {
		log.Fatal(err)
	}
	//var configLines []string

	switch {
	case commands.d != "":
		arr := commands.Delimit(configFile)
		fmt.Println(arr)
		//case commands.f:
		//	if !commands.isDelimited {
		//		fmt.Println(string(configFile))
		//		configLines = strings.Split(string(configFile), "\t") // разделяем каждое слово строки на элементы массива
		//		fmt.Println(configLines[0])
		//	}
		//	Fields(configLines)
		//case commands.s:
		//	fmt.Println("Only-delimited...")
		//
		//	configLines = strings.Split(string(configFile), "\n") // разделяем каждое слово строки на элементы массива
		//	if !commands.isDelimited {
		//		commands.delimiter = "\t"
		//	}
		//	for s := range configLines {
		//		if strings.Contains(configLines[s], commands.delimiter) {
		//			fmt.Println(configLines[s])
		//		}
		//	}
	}

}
