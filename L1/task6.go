package main

import (
	"fmt"
	"sync"
)

//Реализовать все возможные способы остановки выполнения горутины.
// Где взять все способы ???!??!??
func main() {
	var wg sync.WaitGroup

	StopChan := make(chan bool)

	wg.Add(1)

	go func() {
		for {
			select {
			case <-StopChan:
				return
			default:
				defer wg.Done()
				fmt.Println("HI")
			}
		}
	}()

	wg.Wait()
	StopChan <- true

}
