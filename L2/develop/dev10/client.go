package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=7s 127.0.0.1 8081

go-telnet 127.0.0.1 8081
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {

	var wg sync.WaitGroup //переменная для синхронизации горутин

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	if len(text) == 0 {
		fmt.Println("Input error.")
		return
	}

	commands := strings.Split(text, " ")

	var timeout int
	if strings.Contains(text, "--timeout") {
		time := strings.Split(commands[1], "=")
		time = strings.Split(time[1], "s")
		timeout, _ = strconv.Atoi(time[0])
	} else {
		timeout = 10
	}

	fmt.Println(timeout)

	host := commands[len(commands)-2]
	port := commands[len(commands)-1]

	connectTo := host + ":" + port

	d := net.Dialer{
		Timeout: time.Duration(timeout) * time.Second,
	}

	conn, err := d.Dial("tcp", connectTo) //// Подключаемся к сокету

	errChan := make(chan error)
	shutdownCh := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-errChan:
				shutdownCh <- struct{}{}
			}
		}
	}()

	ctx := context.TODO()
	context, cancelCtx := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancelCtx()

	for {
		select {
		case <-context.Done():
			fmt.Println("Timout has finished")
			return
		case <-shutdownCh:
			fmt.Println("Break by Ctrl+D")
			return
		default:
			if err != nil {
				break
			}

			reader := bufio.NewReader(os.Stdin) // Чтение входных данных от stdin
			fmt.Print("Text to send: ")
			text, errorOfSocket := reader.ReadString('\n') // Отправляем в socket

			if errorOfSocket == io.EOF {
				errChan <- errorOfSocket
			} else {
				fmt.Fprintf(conn, text+"\n") // Прослушиваем ответ

				message, ok := bufio.NewReader(conn).ReadString('\n')

				if ok != nil {
					fmt.Println("Server has stopped.")
					return
				}
				fmt.Print("Message from server: " + message)
			}

		}

	}

	wg.Wait()
	conn.Close()
}
