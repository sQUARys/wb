package main

import (
	"fmt"
	"sync"
)

//Разработать конвейер чисел.
//Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout.

func main() {
	var wg sync.WaitGroup
	var count int = 0
	const ArrLen = 5

	x := [ArrLen]int{2, 5, 10, 60, 23}

	FirstChan := make(chan int, 100)
	SecondChan := make(chan int, 100)

	wg.Add(1)
	go sendToFirstCH(x, FirstChan, &wg)

	for {
		select {
		case x := <-FirstChan:
			SecondChan <- (x * 2)
		case in := <-SecondChan:
			fmt.Println(in)
			count++
		default:
			if count == 5 {
				return
			}
		}
	}

	wg.Wait()

}

func sendToFirstCH(x [5]int, firstch chan int, wg *sync.WaitGroup) {

	for i := 0; i < 5; i++ {
		firstch <- x[i]
	}

	defer wg.Done()

}
