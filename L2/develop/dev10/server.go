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
		message, _ := bufio.NewReader(conn).ReadString('\n') // Будем прослушивать все сообщения разделенные \n
		fmt.Print("Message Received:", message)              // Распечатываем полученое сообщение
		newMessage := strings.ToUpper(message)               // Процесс выборки для полученной строки
		conn.Write([]byte(newMessage + "\n"))                // Отправить новую строку обратно клиенту

	}
}
