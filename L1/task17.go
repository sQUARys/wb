package main

//Реализовать бинарный поиск встроенными методами языка.

import (
	"fmt"
	"sort"
)

func BinarySearch(array []int, search_number int) (result bool) {

	lowKey := 0               // первый индекс
	highKey := len(array) - 1 // последний индекс
	if (array[lowKey] > search_number) || (array[highKey] < search_number) {
		return // нужное значение не в диапазоне данных
	}
	for lowKey <= highKey {
		// уменьшаем список рекурсивно
		mid := (lowKey + highKey) / 2 // середина
		if array[mid] == search_number {
			result = true // мы нашли значение
			fmt.Println(mid)
			return
		}
		if array[mid] < search_number {
			// если поиск больше середины - мы берем только блок с большими значениями увеличивая lowKey
			lowKey = mid + 1
			continue
		}
		if array[mid] > search_number {
			// если поиск меньше середины - мы берем блок с меньшими значениями уменьшая highKey
			highKey = mid - 1
		}
	}
	return
}

func main() {
	a := []int{2, 4, 2, 0, 1}
	sort.Ints(a)

	BinarySearch(a, 2)
}
