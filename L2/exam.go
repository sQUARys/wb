package main

import (
	"fmt"
	"io"
	"net"
)

type Server struct {
	connection        net.Conn
	errorOfConnection error
	connectToSite     string
}

func (srv *Server) hasConnection() {
	str := "Hi"
	srv.connectToSite = str
}

func main() {
	err := io.EOF
	server := Server{
		errorOfConnection: err,
	}
	fmt.Println(server)
	server.hasConnection()
	fmt.Println(server)
}
