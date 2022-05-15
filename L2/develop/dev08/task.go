package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps ps not done ??
поддержать fork/exec команды not done ??
конвеер на пайпах done

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Buffer struct {
	data []string
}

func pwd() string {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
	}
	return currentDir
}

func cd(path string) {
	os.Chdir(path)
}

func echo(str string) string {
	cmd := exec.Command("echo", str)
	strOut, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(strOut)
}

func kill() {
	// Start a process:
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
		err := cmd.Process.Kill()
		if err != nil {
			log.Fatal("Failed to kill process: ", err)
		}
		log.Println("Process killed as timeout reached")
	case err := <-doneCh:
		if err != nil {
			log.Fatalf("Process finished with error = %v", err)
		}
		log.Print("Process finished successfully")
	}
}

func main() {

	fmt.Print("Введите команду(можно через конвеер на пайпах):")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	if len(text) == 0 {
		fmt.Println("Input error.")
		return
	}

	buf := Buffer{}

	commands := strings.Split(text, "|")

	isTouchedBuf := false

	for i := 0; i < len(commands); i++ {
		if strings.Contains(commands[i], "cd") {
			path := strings.Split(commands[i], " ")
			cd(path[1])
		}
		if strings.Contains(commands[i], "pwd") {
			buf.data = append(buf.data, pwd())
			isTouchedBuf = true
		}
		if strings.Contains(commands[i], "echo") {
			echoStr := strings.Split(commands[i], "'")[1]
			if isTouchedBuf {
				echoArr := strings.Join(buf.data, "")
				buf.data = append(buf.data, echo(echoArr))
				buf.data = append(buf.data, echo(echoStr))

			} else {
				buf.data = append(buf.data, echo(echoStr))
				isTouchedBuf = true
			}
		}
	}

	fmt.Println(strings.Join(buf.data, " "))

	//fmt.Println(pwd())
	//cd("..")
	//fmt.Println(pwd())
	//fmt.Println(echo("Hello world", "How are you"))

	//PS - ?
	//matches, _ := filepath.Glob("/proc/*/exe")
	//for _, file := range matches {
	//	target, _ := os.Readlink(file)
	//	fmt.Println(filepath.Base(target))
	//
	//	if len(target) > 0 {
	//		fmt.Printf("%+v\n", target)
	//	}
	//}

	//kill()
}
