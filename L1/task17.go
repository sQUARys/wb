package main

//Реализовать бинарный поиск встроенными методами языка.

import (
	"fmt"
	"sort"
)

func BinarySearch(array []int, SearchNumber int) (result bool) { // функция бинарного поиска

	left := 0               // первый индекс массива
	right := len(array) - 1 // последний индекс массива

	if (array[left] > SearchNumber) || (array[right] < SearchNumber) { // значение вышло за границы массива
		return // выходим из функции
	}

	for left <= right { // пока первый индекс меньше либо равен последнего
		mid := (left + right) / 2       // серединный индекс массива
		if array[mid] == SearchNumber { // если мы нашли значение
			result = true
			fmt.Println(mid) // выводим что значение найдено
			return           // выходим из функции
		}
		if array[mid] < SearchNumber { // если центральное значение меньше искомого значения
			left = mid + 1 // левую границу смещаем в середину
		}
		if array[mid] > SearchNumber { // если центральное значение больше искомого значения
			right = mid - 1 // правую границу смещаем в середину
		}
	}
	return
}

func main() {
	a := []int{2, 4, 2, 0, 1} // неотсортированный массив
	sort.Ints(a)              // сортируем его

	BinarySearch(a, 2) // ищем значение 2 в массиве а
}
