package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

func Sleep(duration int) { // собственная функция sleep
	dur := time.After(time.Second * time.Duration(duration)) //канал, в который по истечении времени After передает сообщение о завершении времени
	for {
		select {
		case <-dur: // если в канал поступило сообщение об истечении времени
			fmt.Println("Stop sleeping") // завершаем работу цикла
			return
		default:
			fmt.Println("Waiting") // иначе выводим сообщение об ожидании
		}
	}
}

func main() {
	fmt.Println("Before...")
	Sleep(2) //ждем 2 секунды
	fmt.Println("After...")
}
