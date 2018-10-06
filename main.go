package main

import (
	"echoServer/echoserver"
	"fmt"
)

const (
	IP   = "127.0.0.1"
	PORT = "12345"
)

func main() {
	fmt.Println("Server run on", IP+":"+PORT)
	server := &echoserver.Server{Ip: IP, Port: PORT}
	server.Run()
	return
}
