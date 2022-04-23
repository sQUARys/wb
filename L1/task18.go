package main

import (
	"fmt"
	"sync"
)

//Реализовать структуру-счетчик,
//которая будет инкрементироваться в конкурентной среде.
//	По завершению программа должна выводить итоговое значение счетчика.

type Counter struct { // структура-счетчик
	Number int
}

func main() {
	var wg sync.WaitGroup //переменная для синхронизации горутин
	var mut sync.Mutex    // определяем мьютекс

	num := Counter{ // создаем структуру
		Number: 0, // задаем полю структуры начальное значение
	}

	for i := 0; i < 100000; i++ {
		wg.Add(1) // добавляем к синхронизации еще одну горутину
		go func(m *sync.Mutex) { // запускаем горутину
			m.Lock()        // блокируем num для других горутин
			num.Number++    // добавляем единицу к полю Number
			m.Unlock()      // разблокировываем num для других горутин
			defer wg.Done() // уведомляем о завершении работы горутины
		}(&mut)
	}

	wg.Wait() // ждем завершения всех горутин

	fmt.Println(num.Number) // выводим значение

}