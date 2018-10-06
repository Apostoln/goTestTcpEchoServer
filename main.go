package main

import (
	"echoServer/echoserver"
	"flag"
	"fmt"
)

var (
	ipFlag   string
	portFlag string
)

func init() {
	flag.StringVar(&ipFlag, "ip", "", "server ipFlag address")
	flag.StringVar(&portFlag, "port", "", "server portFlag")
	flag.Parse()
}

func main() {
	fmt.Println("Server run on", ipFlag+":"+portFlag)
	server := echoserver.New(ipFlag, portFlag)
	server.Run()
	return
}
