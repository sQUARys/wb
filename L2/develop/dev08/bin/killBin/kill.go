package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("sleep", "5")

	err := cmd.Start()

	if err != nil {
		fmt.Println(err.Error())
	}
	// Wait for the process to finish or kill it after a timeout (whichever happens first):
	doneCh := make(chan error, 1)

	go func() {
		doneCh <- cmd.Wait()
	}()

	select {
	case <-time.After(7 * time.Second):
		errKill := cmd.Process.Kill()
		if errKill != nil {
			fmt.Print("Failed to kill process: ", errKill)
		}
		fmt.Println("Process killed as timeout reached")
	case e := <-doneCh:
		if e != nil {
			fmt.Print("Process finished with error = %v", e)
		}
		fmt.Print("Process finished successfully(Killed)")
	}
}
