package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

type Ages struct {
	Id  int
	Age int
}

func New(id int, age int) Ages {
	return Ages{
		Id:  id,
		Age: age,
	}
}

func main() {
	var N int
	var wg sync.WaitGroup
	//var mut sync.Mutex

	fmt.Scanf("%d\n", &N)

	ages := make(chan Ages, 100)
	SignalChan := make(chan os.Signal, 1)
	DoneChan := make(chan string)

	signal.Notify(SignalChan, os.Interrupt)

	for input := 1; input <= N; input++ {
		str := New(input, input*10)
		ages <- str
	}
	for i := 0; i < N; i++ {
		wg.Add(1)
		go worker(&wg, ages)
	}

	wg.Add(1)
	go func() {
		<-SignalChan
		defer wg.Done()
		fmt.Println("\nReceived an interrupt, stopping services...\n")
		DoneChan <- "Done"
	}()
	<-DoneChan

	close(ages)

	wg.Wait()
}

func worker(wg *sync.WaitGroup, ages <-chan Ages) {
	for str := range ages {
		fmt.Printf("Hello № %d and your  is %d\n", str.Id, str.Age)
	}
	defer wg.Done()
}
