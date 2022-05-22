package main

import (
	"bufio"
	"context"
	"flag"
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
go-telnet -timeout=15s 127.0.0.1 8081

go-telnet 127.0.0.1 8081
go-telnet --timeout=3s 1.1.1.1 123

Command :  go run client.go -timeout=15s 127.0.0.1 8081

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	var mut sync.Mutex
	var wg sync.WaitGroup //переменная для синхронизации горутин

	server := Server{}

	server.setFlags()
	server.onStartProcess(&wg, &mut)
}

type Server struct {
	connection        net.Conn
	errorOfConnection error
	connectToSite     string
	isConnected       bool
	timeout           int
}

func (srv *Server) setFlags() {
	var timeout string
	flag.StringVar(&timeout, "timeout", "10", "A timeout Var")
	flag.Parse()
	args := flag.Args()
	srv.connectToSite = args[0] + ":" + args[1]
	if strings.Contains(timeout, "s") {
		timeout = strings.Split(timeout, "s")[0]
	}
	srv.timeout, _ = strconv.Atoi(timeout)
	fmt.Println(srv.timeout, srv.connectToSite)
}

func (srv *Server) onStartProcess(wg *sync.WaitGroup, mut *sync.Mutex) {

	errChan := make(chan error, 1)
	shutdownCh := make(chan struct{}, 1)

	wg.Add(1)
	go srv.hasConnection(wg, mut)

	wg.Add(1)
	go checkEofError(wg, mut, errChan, shutdownCh)

	ctx := context.Background()
	context, cancelCtx := context.WithTimeout(ctx, time.Duration(srv.timeout)*time.Second)
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
			if srv.isConnected && srv.errorOfConnection == nil {
				reader := bufio.NewReader(os.Stdin) // Чтение входных данных от stdin
				fmt.Print("Text to send: ")
				text, err := reader.ReadString('\n') // Отправляем в socket

				if err != nil {
					errChan <- err
					srv.errorOfConnection = err
					break
				}

				connectionWithServer := srv.connection
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
	defer srv.connection.Close()
	wg.Wait()
}

func (srv *Server) hasConnection(wg *sync.WaitGroup, mut *sync.Mutex) {
	defer wg.Done()
	fmt.Println("Connection...")
	conn, errDial := net.Dial("tcp", srv.connectToSite) //// Подключаемся к сокету
	for errDial != nil {
		conn, errDial = net.Dial("tcp", srv.connectToSite) //// Подключаемся к сокету
	}
	mut.Lock()
	srv.connection = conn
	srv.errorOfConnection = errDial
	srv.isConnected = true
	mut.Unlock()
}

func checkEofError(wg *sync.WaitGroup, mut *sync.Mutex, c chan error, shutdownCh chan struct{}) {
	defer wg.Done()

	isShutdown := false
	for !isShutdown {
		select {
		case err := <-c:
			if err == io.EOF {
				mut.Lock()
				shutdownCh <- struct{}{}
				isShutdown = true
				close(shutdownCh)
				close(c)
				mut.Unlock()
			}
		}
	}
}
