package main

import (
	"fmt"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать.
//По истечению N секунд программа должна завершаться.

func SendName(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}

func main() {
	var N int

	fmt.Scanf("%d\n", &N)

	ids := make(chan int)

	go SendName(ids)
	for {
		select {
		case id := <-ids:
			fmt.Printf("Hello №%d! \n", id)
		case <-time.After(time.Duration(N) * time.Second):
			fmt.Printf("Terminate")
			return
		}
	}
	close(ids)

}
