package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		readBuffer := make([]byte, 1024)
		_, err := conn.Read(readBuffer)
		if err != nil {
			fmt.Println("Unable to read message from connection")
			os.Exit(1)
		}

		_, err = conn.Write([]byte("+PONG\r\n"))
		if err != nil {
			fmt.Println("found an error trying to respond")
			return
		}
	}
}
