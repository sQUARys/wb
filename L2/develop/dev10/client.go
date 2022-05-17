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

//1. Почитай про пакет flag
//2. conn, err := d.Dial("tcp", connectTo) Абсолютно всегда проверяй ошибки DONE
//3. первая горутина в бесконечном цикле. Ты не делаешь в ней break. Дальше, там ты не закрываешь канал в который пишешь DONE
//4. context. - не должен быть на проде. Замени на что-то, например background
//5. в дефолтном кейсе, что за проверка ошибок в самом начале? вообще неясно к чему относится. Ее не должно там быть
//6. conn.Close() как и все другие close делаются через defer

type Server struct {
	connection        net.Conn
	errorOfConnection error
	connectToSite     string
	isConnected       bool
}

func (srv *Server) hasConnection(wg *sync.WaitGroup, mut *sync.Mutex, c chan bool) {
	defer wg.Done()
	fmt.Println("Connection...")
	conn, errDial := net.Dial("tcp", srv.connectToSite) //// Подключаемся к сокету
	for errDial != nil {
		_, errDial = net.Dial("tcp", srv.connectToSite) //// Подключаемся к сокету
	}
	mut.Lock()
	srv.connection = conn
	srv.errorOfConnection = errDial
	srv.isConnected = true
	mut.Unlock()
	c <- true
	close(c)
}

func checkEofError(wg *sync.WaitGroup, c chan error, shutdownCh chan struct{}) {
	defer wg.Done()

	isShutdown := false
	for !isShutdown {
		select {
		case err := <-c:
			if err == io.EOF {
				shutdownCh <- struct{}{}
				isShutdown = true
				close(shutdownCh)
				close(c)
			}
		}
	}
}

func main() {
	var mut sync.Mutex
	var wg sync.WaitGroup //переменная для синхронизации горутин
	server := Server{}

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

	errChan := make(chan error, 1)
	shutdownCh := make(chan struct{}, 1)
	//serverCh := make(chan Server)
	connCh := make(chan bool, 1)

	host := commands[len(commands)-2]
	port := commands[len(commands)-1]

	server.connectToSite = host + ":" + port

	wg.Add(1)
	go server.hasConnection(&wg, &mut, connCh)

	wg.Add(1)
	go checkEofError(&wg, errChan, shutdownCh)

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
		case <-connCh:
			if server.connection != nil {
				reader := bufio.NewReader(os.Stdin) // Чтение входных данных от stdin
				fmt.Print("Text to send: ")
				text, errorOfSocket := reader.ReadString('\n') // Отправляем в socket

				if errorOfSocket != nil {
					errChan <- errorOfSocket
				} else {
					connectionWithServer := server.connection
					fmt.Fprintf(connectionWithServer, text+"\n") // Прослушиваем ответ

					message, ok := bufio.NewReader(connectionWithServer).ReadString('\n')

					if ok != nil {
						fmt.Println("Server has stopped.")
						return
					}
					fmt.Print("Message from server: " + message)
				}
			}

		}
	}

	wg.Wait()
}

//go func() {
//	isShutdown := true
//	defer wg.Done()
//	for isShutdown {
//		select {
//		case <-errChan:
//			shutdownCh <- struct{}{}
//			isShutdown = false
//			close(errChan)
//			close(shutdownCh)
//		}
//	}
//}()

//go func() {
//	isShutdown := true
//	defer wg.Done()
//	for isShutdown {
//		select {
//		case <-errChan:
//			shutdownCh <- struct{}{}
//			isShutdown = false
//			close(errChan)
//			close(shutdownCh)
//		}
//	}
//}()

//go func(srv *Server) { //  ожидание подключение к серверу
//	defer wg.Done()
//
//	var conn net.Conn
//	var errDial error
//	isConnected := false
//
//	for !isConnected {
//		conn, errDial = net.Dial("tcp", connectTo) //// Подключаемся к сокету
//		if errDial == nil {
//			isConnected = true
//			serverStatusCh <- struct{}{}
//			close(serverStatusCh)
//		}
//	}
//	connectCh <- conn
//	close(connectCh)
//}()

//for {
//	select {
//	case <-context.Done():
//		fmt.Println("Timout has finished")
//		return
//	case <-shutdownCh:
//		fmt.Println("Break by Ctrl+D")
//		return
//	case <-serverStatusCh:
//		reader := bufio.NewReader(os.Stdin) // Чтение входных данных от stdin
//		fmt.Print("Text to send: ")
//		text, errorOfSocket := reader.ReadString('\n') // Отправляем в socket
//
//		if errorOfSocket == io.EOF {
//			errChan <- errorOfSocket
//		} else {
//			connectionWithServer := <-connectCh
//			fmt.Fprintf(connectionWithServer, text+"\n") // Прослушиваем ответ
//
//			message, ok := bufio.NewReader(connectionWithServer).ReadString('\n')
//
//			if ok != nil {
//				fmt.Println("Server has stopped.")
//				return
//			}
//			fmt.Print("Message from server: " + message)
//		}
//	}
//}
