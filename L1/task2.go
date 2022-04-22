package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
//и выведет их квадраты в stdout.

func main() {

	var wg sync.WaitGroup
	array := [5]int{2, 4, 6, 8, 10}

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i int) {
			fmt.Println(array[i] * array[i])
			defer wg.Done()
		}(i)
	}

	wg.Wait()

}
