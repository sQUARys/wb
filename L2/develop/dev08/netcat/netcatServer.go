package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	listen     = flag.Bool("l", false, "Listen")
	hostServer = flag.String("h", "localhost", "Host")
	portServer = flag.Int("p", 0, "Port")
)

func main() {
	flag.Parse()
	if *listen {
		startServer()
		return
	}
	if len(flag.Args()) < 2 {
		fmt.Println("Hostname and port required")
		return
	}
}

func startServer() {
	addr := fmt.Sprintf("%s:%d", *hostServer, *portServer)
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Connection...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection from client: %s", err)
		} else {
			go inputReceiver(conn)
		}
	}
}

func inputReceiver(conn net.Conn) {
	_, err := io.Copy(os.Stdout, conn)
	if err != nil {
		fmt.Println(err)
	}
	conn.Close()
}
