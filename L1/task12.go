package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

type Map struct { // структура, хранящая в себе карту
	items map[string]int
}

func onCreateMap() Map { // создание структуры
	items := make(map[string]int) // создание карты

	NewMap := Map{
		items: items, // запись карты в новую структуру
	}
	return NewMap
}

func (m *Map) Set(id int, val string) { // функция для записи в карту по ключу и значению
	m.items[val] = id
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"} // массив для записи в множество

	mapa := onCreateMap() // создаем новую структуру с картой
	for i := 0; i < len(arr); i++ {
		mapa.Set(i, arr[i]) // записываем ключ и значение в карту структуры
	}
	fmt.Println("Your map contain:")
	for key, val := range mapa.items {
		fmt.Printf("	%s with id = %d\n", key, val) // выводим все множество, записанное в карту
	}

}
