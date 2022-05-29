package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps DONE
поддержать fork/exec команды done
конвеер на пайпах done

Реализовать утилиту netcat (nc) клиент DONE
принимать данные из stdin и отправлять в соединение (tcp/udp) DONE
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Buffer struct {
	data []string
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
				fmt.Println(echoArr)
				res := startBinaryFile("echoBin/echo", echoArr)
				buf.data = append(buf.data, res)
			} else {
				if strings.Contains(commands[i], "'") {
					echoStr = strings.Split(commands[i], "'")[1]
				} else {
					echoStr = strings.Split(commands[i], " ")[1]
				}
				res := startBinaryFile("echoBin/echo", strings.TrimSpace(echoStr))
				buf.data = append(buf.data, res)
				isTouchedBuf = true
			}
		}

		if strings.Contains(commands[i], "kill") {
			res := startBinaryFile("killBin/kill", "")
			buf.data = append(buf.data, res)
		}

		if strings.Contains(commands[i], "ps") {
			res := startBinaryFile("psBin/ps", "")
			buf.data = append(buf.data, res)
			isTouchedBuf = true
		}

		if strings.Contains(commands[i], "exec") {
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

		if strings.Contains(commands[i], "fork") {
			res := startBinaryFile("forkBin/fork", "")
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
