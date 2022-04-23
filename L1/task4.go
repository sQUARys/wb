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

type Ages struct { // структура хранящая в себе id и возсраст
	Id  int
	Age int
}

func New(id int, age int) Ages { // функция создающая новую структуру, которая принимает значение id и значение возраста
	return Ages{
		Id:  id,
		Age: age,
	}
}

func main() {
	var N int             // переменная для ввода пользователем количества воркеров
	var wg sync.WaitGroup //переменная для синхронизации горутин

	fmt.Scanf("%d\n", &N) // считываение пользовательского ввода

	ages := make(chan Ages, 100)          // канал для передачи структуры Ages
	SignalChan := make(chan os.Signal, 1) // канал для считывания команды Ctrl+C
	DoneChan := make(chan string)         // канал для всей информации о завершении программы

	signal.Notify(SignalChan, os.Interrupt) //регистрирует в канал SignalChan состояние символа Ctrl+C

	for input := 1; input <= N; input++ {
		str := New(input, input*10)
		ages <- str // запись в канал структуры с id = input и age = input * 10
	}

	for i := 0; i < N; i++ {
		wg.Add(1)            // в группе теперь +1 горутина
		go worker(&wg, ages) // запуск горутины, в которую передали ссылку на переменную синхронизации и канал структур
	}

	wg.Add(1)   // в группе теперь +1 горутина
	go func() { //запуск горутины
		<-SignalChan    // ждем, когда в канал поступит уведомление о нажатии Ctrl+C
		defer wg.Done() // ждем пока все горутины не завершат работу
		fmt.Println("\nReceived an interrupt, stopping services...\n")
		DoneChan <- "Done" // передаем в канал завершения программы, что необходимо завершить программу
	}()
	<-DoneChan // ждем, когда в канал поступит сообщение о завершении программы

	wg.Wait()   //ждем выполнения всех горутин
	close(ages) // закрываем каналы
	close(DoneChan)
	close(SignalChan)
}

func worker(wg *sync.WaitGroup, ages <-chan Ages) { // воркер
	for str := range ages { // пока в канале есть какая-либо структура, считываем
		fmt.Printf("Hello № %d and your  is %d\n", str.Id, str.Age)
	}
	defer wg.Done() // горутина закончила работу
}
