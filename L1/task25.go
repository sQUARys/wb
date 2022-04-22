package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

func myOwnSleep(duration int) {
	dur := time.After(time.Second * time.Duration(duration))
	for {
		select {
		case <-dur:
			fmt.Println("Stop sleeping")
			return
		default:
			fmt.Println("Waiting")
		}
	}
}

func main() {
	fmt.Println("Before...")
	myOwnSleep(2)
	fmt.Println("After...")
}
