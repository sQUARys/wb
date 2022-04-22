package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {
	array := [5]int{2, 4, 6, 8, 10}

	var sum int
	var wg sync.WaitGroup
	var mut sync.Mutex

	for i := 0; i < len(array); i++ {
		wg.Add(1)

		go func(i int, m *sync.Mutex) {
			m.Lock()
			sum += array[i] * array[i]
			m.Unlock()
			defer wg.Done()
		}(i, &mut)
	}

	wg.Wait()
	fmt.Println("RESULT : ", sum)
}
