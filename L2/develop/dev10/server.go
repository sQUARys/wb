package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Launching server...")

	// Устанавливаем прослушивание порта
	ln, _ := net.Listen("tcp", "127.0.0.1:8081")

	// Открываем порт
	conn, _ := ln.Accept()

	for {
		message, err := bufio.NewReader(conn).ReadString('\n') // Будем прослушивать все сообщения разделенные \n
		message = strings.Trim(message, "\n")
		isCompare := strings.Compare(message, "stop")

		if isCompare == 0 || err != nil {
			fmt.Println("Server has interrupted")
			break
		}
		newMessage := strings.ToUpper(message)    // Процесс выборки для полученной строки
		conn.Write([]byte(newMessage + "\n"))     // Отправить новую строку обратно клиенту
		fmt.Println("Message Received:", message) // Распечатываем полученое сообщение

	}

}
