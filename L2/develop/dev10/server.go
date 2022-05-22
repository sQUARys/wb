package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	fmt.Println("Launching server...")

	// Устанавливаем прослушивание порта
	ln, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatalln(err)
	}

	// Открываем порт
	conn, err := ln.Accept()
	if err != nil {
		log.Fatalln(err)
	}
	for {
		message, err := bufio.NewReader(conn).ReadString('\n') // Будем прослушивать все сообщения разделенные \n
		if err != nil {
			fmt.Println("Server has interrupted")
			break
		}

		message = strings.Trim(message, "\n")
		isCompare := strings.Compare(message, "stop")

		if isCompare == 0 {
			fmt.Println("Server has interrupted")
			break
		}
		newMessage := strings.ToUpper(message)    // Процесс выборки для полученной строки
		conn.Write([]byte(newMessage + "\n"))     // Отправить новую строку обратно клиенту
		fmt.Println("Message Received:", message) // Распечатываем полученое сообщение

	}

}
