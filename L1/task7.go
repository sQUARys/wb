package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map.

type DefMap struct { // структура хранящая в себе мапу и мьютекс для ее блокировки и разблокировки
	mx sync.Mutex
	m  map[int]string
}

func New() *DefMap { // функция создания новой структуры DefMap с пустой мапой
	return &DefMap{
		m: make(map[int]string),
	}
}

func (dm *DefMap) Set(id int, value string) { // функция установки id и value в мапу структуры DefMap
	dm.mx.Lock()     // блокируем доступ к структуре для других функций для более безопасной работы
	dm.m[id] = value // записываем в структуру значение по ключу
	dm.mx.Unlock()   // разблокируем достук к структуре
}

func (dm *DefMap) Get(id int) (string, bool) { // функция получения по id значения мапы с соответсвенным ключом
	dm.mx.Lock()        // блокируем доступ к структуре для других функций для более безопасной работы
	val, ok := dm.m[id] // записываем в структуру значение по ключу
	dm.mx.Unlock()      // разблокируем доступ к структуре

	return val, ok
}

func main() {
	var wg sync.WaitGroup                                                  //переменная для синхронизации горутин
	lines := [5]string{"Dmitrii", "Nikolay", "Rostislav", "Elena", "Oleg"} // массив имен

	defmap := New() // создаем новую структуру

	for i := 0; i < 5; i++ {
		wg.Add(1)        // добавляем к синхронизации горутину
		go func(i int) { // запуск горутины
			defmap.Set(i, lines[i]) // записываем ключ и значения в структуру
			defer wg.Done()         // горутина закончила работу
		}(i)
	}

	wg.Wait() // ждем завершения всех горутин

	fmt.Println("Waiting is over.")
	fmt.Println("Your map:")
	for i := 0; i < 5; i++ {
		val, ok := defmap.Get(i) // считываем значения по ключу из мапы
		if ok {                  // если значение успешно получено, выводим
			fmt.Printf("Your key: %d. Your value: %s.\n", i, val)
		}
	}

}
