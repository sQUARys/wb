package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	pathOfStartedBin, _ := r.ReadString('\n')

	binary, lookErr := exec.LookPath(pathOfStartedBin)
	if lookErr != nil {
		log.Fatal(lookErr)
		return
	}

	args := []string{pathOfStartedBin} //дополнительные команды могут быть: "-a", "-l", "-h"

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		log.Fatal(execErr)
	}
}
