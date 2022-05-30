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

func main() {
	commands := Command{} // создаем пустую структуру
	commands.setFlags()   // считываем команды пользователя

	configFile, err := ioutil.ReadFile("data1.txt") // считываем данные из файла

	if err != nil {
		log.Fatal(err)
	}
	var configLines []string

	if commands.d != "" { // если пользователь ввел команду d
		configLines = commands.Delimit(configFile)
		fmt.Println(configLines)
	}
	if commands.f != 0 { // если пользователь ввел команду f
		ret := commands.Fields(configLines)
		fmt.Println(ret)
	}
	if commands.s { //если пользователь ввел команду s
		ret := commands.Separated(configFile)
		fmt.Println(ret)
	}

}

type Command struct {
	d string
	f int
	s bool
}

func (c *Command) setFlags() {
	flag.StringVar(&c.d, "d", "\t", "delimiter")
	flag.IntVar(&c.f, "f", 0, "fields")
	flag.BoolVar(&c.s, "s", false, "separated")
	flag.Parse()
}

func (c *Command) Delimit(configFile []byte) []string {
	fmt.Println("Delimiter...")
	configLines := strings.Split(string(configFile), c.d) // разделяем каждое слово строки на элементы массива
	return configLines
}

func (c *Command) Fields(configLines []string) []string {
	fmt.Println("Fields...")
	var fields []string
	var result []string

	for f := range configLines { // проходим по массиву
		fields = strings.Fields(configLines[f]) // разделяем по разделителю
		if c.f > len(fields) {
			fmt.Println("Вы ввели поле, которое выходит за границы строк. Введите пожалуйста значение поменьше")
			return []string{}
		}
		result = append(result, fields[c.f-1]) // выводим
	}
	return result
}

func (c *Command) Separated(configFile []byte) []string {
	fmt.Println("Only-delimited...")

	result := []string{}

	configLines := strings.Split(string(configFile), "\n") // разделяем каждое слово строки на элементы массива
	for s := range configLines {                           // проходим по массиву
		if strings.Contains(configLines[s], c.d) { // если в строке есть разделитель выводим ее
			result = append(result, configLines[s])
		}
	}
	return result
}
