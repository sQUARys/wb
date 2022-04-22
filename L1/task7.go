package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map.

type DefMap struct {
	mx sync.Mutex
	m  map[int]string
}

func New() *DefMap {
	return &DefMap{
		m: make(map[int]string),
	}
}

func (dm *DefMap) Set(id int, value string) {
	dm.mx.Lock()
	dm.m[id] = value
	dm.mx.Unlock()
}

func (dm *DefMap) Get(id int) (string, bool) {
	dm.mx.Lock()
	val, ok := dm.m[id]
	dm.mx.Unlock()

	return val, ok
}

func main() {
	var wg sync.WaitGroup
	lines := [5]string{"Dmitrii", "Nikolay", "Rostislav", "Elena", "Oleg"}

	defmap := New()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defmap.Set(i, lines[i])
			defer wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("Waiting is over.")
	fmt.Println("Your map:")
	for i := 0; i < 5; i++ {
		val, ok := defmap.Get(i)
		if ok {
			fmt.Printf("Your key: %d. Your value: %s.\n", i, val)
		}
	}

}
