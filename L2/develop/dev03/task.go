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

//1. fmt.Println("Вы выбрали отсортировать по числовому значению.") - это никому не надо. Просто выполняй результат также, как и рабтает утилита sort линукса. man sort - нужно для того, чтобы понять, как она это делает. Реализовывать man не надо
//2. Не пиши ты функции выше мейна. Ты не в си. Там так делается не потому, что это хорошо, а потому что по-другому не скомпилируется. Это неудобно так делать. Я (и все остальные) читаем код сверху вниз, поэтому мейн первый. Не надо функции делать экспортируемыми. У тебя внутри мейна должна вызываться одна функция sort с аргументами какими-нибудь, а уже внутри любые другие неэкспорируемые функции.
//3. val1, _ := strconv.Atoi(sl[i]) - без err это просто бан

func main() {
	str := []string{"", "4", "", "1", "2"}

	commands := Command{}
	commands.flagSet()
	commands.input = str

	commands.sort()
}

func (c Command) sort() {
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

type Command struct {
	input []string
	k     int
	n     bool
	r     bool
	u     bool
}

func (c *Command) flagSet() {
	flag.IntVar(&c.k, "k", -1, "Column")
	flag.BoolVar(&c.n, "n", false, "sorting by int value")
	flag.BoolVar(&c.r, "r", false, "reverse sorting")
	flag.BoolVar(&c.u, "u", false, "sorting without repeated")

	flag.Parse()
}

func reverse(sl []string) []string {
	sort.Strings(sl)
	sort.Sort(sort.Reverse(sort.StringSlice(sl)))
	return sl
}

func sortWithoutRepeat(sl []string) []string {
	var mem []string
	sort.Strings(sl)
	for i := range sl {
		if len(mem) == sort.SearchStrings(mem, sl[i]) { // если не найдено такой строки
			mem = append(mem, sl[i])
		}
	}
	return mem
}

func sortBySpecialColumn(sl []string, k int) []string {
	sort.Slice(sl, func(i, j int) bool {
		if k >= len(sl[i]) || k >= len(sl[j]) {
			return false
		}
		return sl[i][k] < sl[j][k]
	})
	return sl
}

func sortByNumber(sl []string) []string {
	sort.Slice(sl, func(i, j int) bool {
		val1, error := strconv.Atoi(sl[i])
		if error != nil {
			fmt.Println("Error: ", error)
		}
		val2, err := strconv.Atoi(sl[j])
		if err != nil {
			fmt.Println("Error: ", err)
		}
		return val1 < val2
	})
	return sl
}
