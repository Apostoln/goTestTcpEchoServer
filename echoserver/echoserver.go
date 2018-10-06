package echoserver

import (
	"fmt"
	"io"
	"net"
)

const (
	TCP         = "tcp"
	BUFFER_SIZE = 1024
)

type Server struct {
	Ip   string
	Port string
}

func (server *Server) Run() {

	listener, err := net.Listen(TCP, server.Ip+":"+server.Port)
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
