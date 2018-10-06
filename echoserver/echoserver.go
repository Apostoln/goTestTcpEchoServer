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
	Connections map[*net.Conn] bool //set emulation, kill me please
}

func New(ip string, port string) Server {
	res := Server{Ip:ip, Port:port}
	res.Connections = make(map[*net.Conn] bool)
	return res
}

func (server *Server) Write(bytes []byte) {
	for conn, _ := range server.Connections {
		(*conn).Write(bytes) //TODO: how to do it without (*conn)

	}

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

		server.Connections[&connection] = true

		go func() {
			defer connection.Close()
			defer delete(server.Connections, &connection)

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
				server.Write(buffer[:size])
			}
		}()
	}
}
