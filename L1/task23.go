package main

import "fmt"

//Удалить i-ый элемент из слайса.

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 2

	copy(arr[i:], arr[i+1:])
	arr = arr[:(len(arr) - 1)]
	fmt.Println(arr)
}
