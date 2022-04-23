package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Реализовать пересечение двух неупорядоченных множеств.

type Map struct { // структура, хранящая в себе карту
	items map[int]string
}

func onCreateMap() Map { // создание структуры
	items := make(map[int]string) // создание карты

	NewMap := Map{
		items: items, // запись карты
	}
	return NewMap
}

func (m *Map) Set(id int, val string) { // функция для записи в карту по ключу и значению
	m.items[id] = val
}

func (m *Map) Get(id int) string { // получение по ключу значения из карты
	return m.items[id]
}

func main() {
	IdArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}                                                                 // массив id
	ValArr := []string{"Elena", "Maksim", "Roman", "Valery", "Nikita", "Artem", "Dima", "Akim", "Anna", "Nastya"} // массив значений

	First := onCreateMap()  // создаем первую структуру с картой
	Second := onCreateMap() // создаем вторую структуру с картой

	for i := 0; i < len(IdArr); i++ {
		First.Set(IdArr[i], ValArr[i]) // устанавливаем значения в карту первой структуры
	}

	max := 9 // максимальная граница(зависит от длинны массива)
	min := 0 // минимальная граница

	rand.Seed(time.Now().Unix())                         // запуск получения случайного числа
	n := rand.Intn(max-min) + min                        // min ≤ n ≤ max, получения случайного числа и запись его в n
	fmt.Println("Your random value for Second map: ", n) //c n-ого индекса будет осуществлятся запись в карту второй структуры

	for i := n; i < len(IdArr); i++ {
		Second.Set(IdArr[i], ValArr[i]) // устанавливаем с n-ого индекса в карту второй структуры
	}

	fmt.Println("Your maps intersection is : ")
	for FirstKey, FirstValue := range First.items { // сравниваем каждой с каждым значения структур
		for SecondKey, SecondValue := range Second.items {
			if FirstKey == SecondKey && FirstValue == SecondValue { // если нашлись равные ключи и значения
				fmt.Printf("	Key: %d , Value: %s \n", FirstKey, FirstValue) // множества пересеклись и выводим информацию о их перечении
			}
		}
	}
}
