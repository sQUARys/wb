package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
)

/*
=== Утилита sort ==

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки Done
-n — сортировать по числовому значению Done
-r — сортировать в обратном порядке Done
-u — не выводить повторяющиеся строки Done

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	str := []string{"", "4", "", "1", "2"}

	commands := Command{} // создаем пустую структуру
	commands.flagSet()    // считываем из командной строки
	commands.input = str  // вводим наш массив данных

	commands.sort()
}

type Command struct {
	input []string
	k     int
	n     bool
	r     bool
	u     bool
}

func (c *Command) flagSet() { // считывание с командной строки
	flag.IntVar(&c.k, "k", -1, "Column")
	flag.BoolVar(&c.n, "n", false, "sorting by int value")
	flag.BoolVar(&c.r, "r", false, "reverse sorting")
	flag.BoolVar(&c.u, "u", false, "sorting without repeated")

	flag.Parse()
}

func (c Command) sort() { // функция выбора в зависимости от ввода типа сортировка
	if c.n == true {
		fmt.Println(sortByNumber(c.input)) //отсортировать по числовому значению
	}
	if c.k != -1 {
		fmt.Println(sortBySpecialColumn(c.input, c.k)) // отсортировать по определенной колонке
	}
	if c.r == true {
		fmt.Println(reverse(c.input)) // отсортировать в обратном порядке
	}
	if c.u == true {
		fmt.Println(sortWithoutRepeat(c.input)) // отсортировать и не выводить повторяющиеся строки
	}
}

func reverse(sl []string) []string {
	sort.Strings(sl)                              // сортируем по порядку
	sort.Sort(sort.Reverse(sort.StringSlice(sl))) // переворачиваем
	return sl
}

func sortWithoutRepeat(sl []string) []string {
	var mem []string
	sort.Strings(sl) // сортируем маcсив строк
	for i := range sl {
		if len(mem) == sort.SearchStrings(mem, sl[i]) { // если не найдено такой строки
			mem = append(mem, sl[i]) // добавляем уникальное слово
		}
	}
	return mem
}

func sortBySpecialColumn(sl []string, k int) []string {
	sort.Slice(sl, func(i, j int) bool { // функция для сортировки
		if k >= len(sl[i]) || k >= len(sl[j]) { // если длинна меньше введеной колонны не меняем местами
			return false
		}
		return sl[i][k] < sl[j][k]
	})
	return sl
}

func sortByNumber(sl []string) []string {
	sort.Slice(sl, func(i, j int) bool {
		if sl[i] == "" || sl[j] == "" {
			return true
		}
		val1, error := strconv.Atoi(sl[i]) // переводим первую строку в число
		if error != nil {
			fmt.Println("Error: ", error)
			return true
		}
		val2, err := strconv.Atoi(sl[j]) // переводим вторую строку в число
		if err != nil {
			fmt.Println("Error: ", err)
			return true
		}
		return val1 < val2
	})
	return sl
}
