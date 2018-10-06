package main

import (
	"echoServer/echoserver"
	"flag"
	"fmt"
)

var (
	ip   string
	port string
)

func init() {
	flag.StringVar(&ip, "ip", "", "server ip address")
	flag.StringVar(&port, "port", "", "server port")
	flag.Parse()
}

func main() {
	fmt.Println("Server run on", ip+":"+port)
	server := &echoserver.Server{Ip: ip, Port: port}
	server.Run()
	return
}
