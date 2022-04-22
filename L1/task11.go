package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Реализовать пересечение двух неупорядоченных множеств.

type Map struct {
	items map[int]string
}

func onCreateMap() Map {
	items := make(map[int]string)

	NewMap := Map{
		items: items,
	}
	return NewMap
}

func (m *Map) Set(id int, val string) {
	m.items[id] = val
}

func (m *Map) Get(id int) string {
	return m.items[id]
}

func main() {
	IdArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ValArr := []string{"Elena", "Maksim", "Roman", "Valery", "Nikita", "Artem", "Dima", "Akim", "Anna", "Nastya"}

	First := onCreateMap()
	Second := onCreateMap()

	for i := 0; i < len(IdArr); i++ {
		First.Set(IdArr[i], ValArr[i])
	}

	max := 9
	min := 0

	rand.Seed(time.Now().Unix())
	n := rand.Intn(max-min) + min // min ≤ n ≤ max
	fmt.Println("Your random value for Second map: ", n)

	for i := n; i < len(IdArr); i++ {
		Second.Set(IdArr[i], ValArr[i])
	}

	fmt.Println("Your maps intersection is : ")
	for FirstKey, FirstValue := range First.items {
		for SecondKey, SecondValue := range Second.items {
			if FirstKey == SecondKey && FirstValue == SecondValue {
				fmt.Printf("	Key: %d , Value: %s \n", FirstKey, FirstValue)
			}
		}
	}
}
