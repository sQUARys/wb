package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet 127.0.0.1 8081
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	if len(text) == 0 {
		fmt.Println("Input error.")
		return
	}

	commands := strings.Split(text, " ")

	host := commands[len(commands)-2]
	port := commands[len(commands)-1]

	connectTo := host + ":" + port
	conn, _ := net.Dial("tcp", connectTo) //// Подключаемся к сокету
	for {
		reader := bufio.NewReader(os.Stdin) // Чтение входных данных от stdin
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n') // Отправляем в socket
		fmt.Fprintf(conn, text+"\n")       // Прослушиваем ответ
		fmt.Println(conn)

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}

}