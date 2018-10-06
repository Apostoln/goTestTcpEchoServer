package main

import (
	"fmt"
	"io"
	"net"
)

const (
	IP          = "127.0.0.1"
	PORT        = "12345"
	TCP         = "tcp"
	BUFFER_SIZE = 1024
)

func main() {
	fmt.Println("Server run on", IP+":"+PORT)
	listener, err := net.Listen(TCP, IP+":"+PORT)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Error on", connection.RemoteAddr(), err.Error())
			continue
		}

		go func() {
			defer connection.Close()

			buffer := make([]byte, BUFFER_SIZE)
			for {
				size, err := connection.Read(buffer)
				if err != nil {
					if err != io.EOF {
						fmt.Println("Error on", connection.RemoteAddr(), err.Error())
					}
					return
				}
				fmt.Println(connection.RemoteAddr(), ">>", string(buffer[:size]))
				connection.Write(buffer[:size])
			}
		}()
	}
}
