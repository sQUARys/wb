package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
//и выведет их квадраты в stdout.

func main() {

	var wg sync.WaitGroup //переменная для синхронизации горутин
	var rmut sync.RWMutex

	array := [5]int{2, 4, 6, 8, 10} // массив чисел

	for i := 0; i < 5; i++ {
		wg.Add(1) // в группе теперь +1 горутина

		go func(i int) { // запуск горутины
			rmut.Lock()
			fmt.Println(array[i] * array[i]) // вывод квадрата числа из массива
			rmut.Unlock()
			defer wg.Done() // горутина закончила работу
		}(i)
	}

	wg.Wait() // ждем пока все горутины не завершат работу

}
