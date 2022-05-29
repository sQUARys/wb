package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {

	first := 4
	second := 10
	ret, pid, err := syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0) // ret - parents pid, pid - child pid

	if err != 0 {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	if pid > 0 {
		second++
		fmt.Println("In parent:", ret, first, second)
	} else { //pid == 0
		first++
		fmt.Println("In child:", pid, first, second)
	}
}
