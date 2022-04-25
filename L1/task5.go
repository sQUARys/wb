package main

import (
	"fmt"
	"sync"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать.
//По истечению N секунд программа должна завершаться.

func SendName(c chan int, wg *sync.WaitGroup) { // функция которая записывает в канал
	defer wg.Done()

	for i := 0; i < 10; i++ {
		c <- i // записываем значение i в канал
	}
	close(c)
}

func main() {
	var N int // переменная пользовательского ввода количества секунд работы программы
	var wg sync.WaitGroup

	fmt.Scanf("%d\n", &N) // считываение количества секунд

	ids := make(chan int) // открываем не буферизированный канал

	wg.Add(1)
	go SendName(ids, &wg) // запускаем в горутине функцию

	wg.Add(1)
	go func() {
		defer wg.Done()
		timeCh := time.After(time.Duration(N) * time.Second)
		for {
			select {
			case id, ok := <-ids:
				if ok {
					fmt.Printf("Hello №%d! \n", id) // выводим ее
				}
			case <-timeCh: // если время истекло
				fmt.Printf("Terminate") // выводим сообщение о завершении
				return                  // выходим из цикла
			}
		}
	}()

	wg.Wait()

}

//for {
//	id, ok := <-ids
//	if !ok {
//		<-time.After(time.Duration(N) * time.Second)
//		fmt.Printf("Terminate") // выводим сообщение о завершении
//		return                  // выходим из цикла
//	} else {
//		fmt.Printf("Hello №%d! \n", id) // выводим ее
//	}
//}
