package main

import (
	"fmt"
	"github.com/sQUARys/Go-Modules/task"
)

//Can't find go-modules. Why?
//Error
//go: finding module for package github.com/sQUARys/Go-Modules/task
//go: downloading github.com/sQUARys/Go-Modules v0.0.0-20220512093447-2844b6347ec0
//mod imports
//github.com/sQUARys/Go-Modules/task: module github.com/sQUARys/Go-Modules@latest found (v0.0.0-20220512093447-2844b6347ec0), but does not contain package github.com/sQUARys/Go-Modules/task

func main() {
	message := task.GetTime()
	fmt.Println(message)
}
