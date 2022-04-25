package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

type Map struct { // структура, хранящая в себе карту
	items map[string]struct{}
}

func onCreateMap() Map { // создание структуры
	items := make(map[string]struct{}) // создание карты

	NewMap := Map{
		items: items, // запись карты в новую структуру
	}
	return NewMap
}

func (m *Map) Set(val string) { // функция для записи в карту по ключу и значению
	m.items[val] = struct{}{}
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"} // массив для записи в множество

	mapa := onCreateMap() // создаем новую структуру с картой
	for i := 0; i < len(arr); i++ {
		mapa.Set(arr[i]) // записываем ключ и значение в карту структуры
	}
	fmt.Print("Your map contain:")
	for key, _ := range mapa.items {
		fmt.Printf(" %s  ", key) // выводим все множество, записанное в карту
	}

}
