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

func (c *Command) Fields(configLines []string) {
	fmt.Println("Fields...")
	var fields []string

	if c.f >= len(configLines) {
		fmt.Println("Вы ввели поле, которое выходит за границы строк. Введите пожалуйста значение поменьше")
		return
	}
	for f := range configLines {
		fields = strings.Fields(configLines[f])
		fmt.Println(fields[c.f-1])
	}
}

type Command struct {
	d           string
	f           int
	s           bool
	isDelimited bool
}

func (c *Command) setFlags() {
	flag.StringVar(&c.d, "d", "\t", "delimiter")
	flag.IntVar(&c.f, "f", 0, "fields")
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
	var configLines []string

	if commands.d != "" {
		configLines = commands.Delimit(configFile)
		fmt.Println(configLines)
	}
	if commands.f != 0 {
		commands.Fields(configLines)
	}
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
