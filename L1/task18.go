package main

import (
	"fmt"
	"sync"
)

//Реализовать структуру-счетчик,
//которая будет инкрементироваться в конкурентной среде.
//	По завершению программа должна выводить итоговое значение счетчика.

type Counter struct {
	Number int
}

func main() {
	var wg sync.WaitGroup
	var mut sync.Mutex

	num := Counter{
		Number: 0,
	}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(m *sync.Mutex) {
			m.Lock()
			num.Number++
			m.Unlock()
			defer wg.Done()
		}(&mut)
	}

	wg.Wait()

	fmt.Println(num.Number)

}
