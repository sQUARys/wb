package main

import (
	"fmt"
	"github.com/sQUARys/wb/module/task"
)

func main() {
	message := task.GetTime()
	fmt.Println(message)
}
