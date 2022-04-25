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

	ages := make(chan Ages)               // канал для передачи структуры Ages
	SignalChan := make(chan os.Signal, 1) // канал для считывания команды Ctrl+C
	DoneChan := make(chan string)         // канал для всей информации о завершении программы

	signal.Notify(SignalChan, os.Interrupt) //регистрирует в канал SignalChan состояние символа Ctrl+C

	wg.Add(1)
	go func() {
		defer wg.Done()
		for input := 0; input < N; input++ {
			str := New(input+1, (input+1)*10)
			ages <- str // запись в канал структуры с id = input и age = input * 10
		}
	}()

	for i := 0; i < N; i++ {
		wg.Add(1)            // в группе теперь +1 горутина
		go worker(&wg, ages) // запуск горутины, в которую передали ссылку на переменную синхронизации и канал структур
	}

	wg.Add(1)   // в группе теперь +1 горутина
	go func() { //запуск горутины
		defer wg.Done()
		<-SignalChan // ждем, когда в канал поступит уведомление о нажатии Ctrl+C
		fmt.Println("\nReceived an interrupt, stopping services...\n")
		DoneChan <- "Done" // передаем в канал завершения программы, что необходимо завершить программу
	}()

	<-DoneChan // ждем, когда в канал поступит сообщение о завершении программы

	wg.Wait()
	close(ages) // закрываем каналы
	close(DoneChan)
	close(SignalChan)
}

func worker(wg *sync.WaitGroup, ages <-chan Ages) { // воркер
	defer wg.Done() // горутина закончила работу
	str := <-ages   // пока в канале есть какая-либо структура, считываем
	fmt.Printf("Hello № %d and your  is %d\n", str.Id, str.Age)
}
