package main

import (
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
	//?????
}

func main() {
	var choice string
	str := []string{"abc", "acs", "bfd", "aaa", "aaa", "bbb"}

	fmt.Scanf("%s\n", &choice)

	switch choice {
	case "k":
		var k int
		fmt.Println("Вы выбрали отсортировать по определенной колонке.")
		fmt.Print("Введите индекс колонки:")
		fmt.Scanf("%d\n", &k)
		fmt.Println(SortBySpecialColumn(str, k))
	case "r":
		fmt.Println("Вы выбрали отсортировать в обратном порядке.")
		fmt.Println(Reverse(str))
	case "u":
		fmt.Println("Вы выбрали отсортировать и не выводить повторяющиеся строки.")
		fmt.Println(SortWithoutRepeat(str))
	default:
		fmt.Println("Вы ввели несуществующую команду")
	}

}
