package main

import (
	"fmt"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать.
//По истечению N секунд программа должна завершаться.

func SendName(c chan int) { // функция которая записывает в канал
	for i := 0; i < 10; i++ {
		c <- i // записываем значение i в канал
	}
}

func main() {
	var N int // переменная пользовательского ввода количества секунд работы программы

	fmt.Scanf("%d\n", &N) // считываение количества секунд

	ids := make(chan int) // открываем не буферизированный канал

	go SendName(ids) // запускаем в горутине функцию
	for {
		select {
		case id := <-ids: // если в канал поступила информация
			fmt.Printf("Hello №%d! \n", id) // выводим ее
		case <-time.After(time.Duration(N) * time.Second): // если время истекло
			fmt.Printf("Terminate") // выводим сообщение о завершении
			return                  // выходим из цикла
		}
	}
	close(ids) //закрываем канал

}