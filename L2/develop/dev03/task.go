package main

import (
	"flag"
	"fmt"
	"sort"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки Done
-n — сортировать по числовому значению ??
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

func Reverse(sl []string) []string {
	sort.Strings(sl)
	sort.Sort(sort.Reverse(sort.StringSlice(sl)))
	return sl
}

func SortWithoutRepeat(sl []string) []string {
	var mem []string
	sort.Strings(sl)
	for i := range sl {
		if len(mem) == sort.SearchStrings(mem, sl[i]) { // если не найдено такой строки
			mem = append(mem, sl[i])
		}
	}
	return mem
}

func SortBySpecialColumn(sl []string, k int) []string {
	sort.Slice(sl, func(i, j int) bool {
		return sl[i][k] < sl[j][k]
	})
	return sl
}

func SortByNumber(sl []string) {
	//????? что значит сортировать по числовому значению
}

type Command struct {
	k int
	n bool
	r bool
	u bool
}

func (c *Command) flagSet() {
	flag.IntVar(&c.k, "k", -1, "Column")
	flag.BoolVar(&c.n, "n", false, "sorting by int value")
	flag.BoolVar(&c.r, "r", false, "reverse sorting")
	flag.BoolVar(&c.u, "u", false, "sorting without repeated")

	flag.Parse()
}

func main() {
	str := []string{"abc", "acs", "bfd", "aaa", "aaa", "bbb"}

	commands := Command{}
	defaultCommand := Command{
		k: -1,
		n: false,
		r: false,
		u: false,
	}
	commands.flagSet()
	fmt.Println(commands.k)

	if commands.n == true {

	}
	if commands.k != -1 {
		fmt.Println("Вы выбрали отсортировать по определенной колонке.")
		fmt.Println(SortBySpecialColumn(str, commands.k))
	}
	if commands.r == true {
		fmt.Println("Вы выбрали отсортировать в обратном порядке.")
		fmt.Println(Reverse(str))
	}
	if commands.u == true {
		fmt.Println("Вы выбрали отсортировать и не выводить повторяющиеся строки.")
		fmt.Println(SortWithoutRepeat(str))
	}

	if defaultCommand == commands {
		fmt.Println("Вы ввели несуществующую команду")
	}
}
