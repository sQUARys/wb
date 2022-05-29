package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps DONE
поддержать fork/execBin команды not done доделать
конвеер на пайпах done

Реализовать утилиту netcat (nc) клиент DONE
принимать данные из stdin и отправлять в соединение (tcp/udp) DONE
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

func echo(str string) string {
	cmd := exec.Command("echo", str)
	strOut, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(strOut)
}

func kill() {
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

func ps() {
	PS, err := exec.LookPath("ps")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(PS)
	command := []string{"ps", "-a", "-x"}
	env := os.Environ()
	err = syscall.Exec(PS, command, env)
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

	var echoStr string

	for i := 0; i < len(commands); i++ {
		if strings.Contains(commands[i], "cd") {
			path := strings.Split(commands[i], " ")
			if len(path) == 1 {
				log.Fatal("Your cd command don't have an argument")
			}
			res := startBinaryFile("cdBin/cd", path[1])
			buf.data = append(buf.data, res)
		}

		if strings.Contains(commands[i], "pwd") {
			res := startBinaryFile("pwdBin/pwd", "")
			buf.data = append(buf.data, res)
			isTouchedBuf = true
		}

		if strings.Contains(commands[i], "echo") {
			if isTouchedBuf {
				echoArr := strings.Join(buf.data, "")
				buf.data = append(buf.data, echo(echoArr))
			} else {
				if strings.Contains(commands[i], "'") {
					echoStr = strings.Split(commands[i], "'")[1]
				}
				fmt.Println(echo(echoStr))
				buf.data = append(buf.data, echo(echoStr))
				isTouchedBuf = true
			}
		}
		if strings.Contains(commands[i], "kill") {
			kill()
		}
		if strings.Contains(commands[i], "ps") {
			ps()
		}
		if strings.Contains(commands[i], "execBin") {
			var pathToBin string
			if isTouchedBuf {
				pathToBin = strings.Join(buf.data, "")
			} else {
				pathToBin = "ls" // дефолт
				isTouchedBuf = true
			}

			res := startBinaryFile("execBin/exec", strings.TrimSpace(pathToBin))
			buf.data = append(buf.data, res)
		}
	}

	fmt.Println(strings.Join(buf.data, " "))
}

func startBinaryFile(binPath string, data string) string {
	defaultPath := "/Users/roman/Desktop/Work/WorkRepo/L2/develop/dev08/bin/"
	cmd := exec.Command(defaultPath + binPath)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, data)
	}()

	out, errOut := cmd.CombinedOutput()
	if errOut != nil {
		log.Fatal(errOut)
	}

	cmd.Wait()
	return string(out)

}
