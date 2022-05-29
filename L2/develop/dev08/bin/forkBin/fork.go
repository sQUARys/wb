package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	foo := 4
	bar := 10
	ret, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)

	if err != 0 {
		os.Exit(2)
	}

	if ret > 0 {
		bar++
		fmt.Println("In parent:", ret, foo, bar)
		return
	}

	foo++
	fmt.Println("In child:", ret, foo, bar)
}
