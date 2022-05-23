package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var (
	host = flag.String("h", "localhost", "Host")
	port = flag.Int("p", 0, "Port")
)

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		fmt.Printf("Can't connect to server: %s\n", err)
		return
	}
	_, err = io.Copy(conn, os.Stdin)
	if err != nil {
		fmt.Printf("Connection error: %s\n", err)
	}
}
