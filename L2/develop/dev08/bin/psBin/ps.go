package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	PS, err := exec.LookPath("ps")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(PS)
	command := []string{"ps"}
	env := os.Environ()
	err = syscall.Exec(PS, command, env)
}
