package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

type Map struct {
	items map[string]int
}

func onCreateMap() Map {
	items := make(map[string]int)

	NewMap := Map{
		items: items,
	}
	return NewMap
}

func (m *Map) Set(id int, val string) {
	m.items[val] = id
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	mapa := onCreateMap()
	for i := 0; i < len(arr); i++ {
		mapa.Set(i, arr[i])
	}
	fmt.Println("Your map contain:")
	for key, val := range mapa.items {
		fmt.Printf("	%s with id = %d\n", key, val)
	}

}
