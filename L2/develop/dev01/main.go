package main

import (
	"fmt"
	task "github.com/sQUARys/Go-Modules"
)

func main() {
	message := task.GetTime() // вызываем функция из нашего модуля
	fmt.Println(message)      // выводим результат
}
