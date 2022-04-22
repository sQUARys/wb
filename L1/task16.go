package main

import (
	"fmt"
)

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := len(a) - 1

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func main() {
	a := []int{2, 4, 2, 0, 1}
	fmt.Println(quicksort(a))
}
