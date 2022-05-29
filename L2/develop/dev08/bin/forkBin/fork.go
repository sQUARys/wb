package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {

	foo := 4
	bar := 10
	ret, pid, err := syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0) // ret - parents pid, pid - child pid

	if err != 0 {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	if pid > 0 {
		bar++
		fmt.Println("In parent:", ret, foo, bar)
	} else { //pid == 0
		foo++
		fmt.Println("In child:", pid, foo, bar)
	}
}
