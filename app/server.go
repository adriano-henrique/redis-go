package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	handleConnection(conn)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	readBuffer := make([]byte, 1024)
	_, err := conn.Read(readBuffer)
	if err != nil {
		fmt.Println("Unable to read message from connection")
		os.Exit(1)
	}
	requestContent := string(readBuffer)
	var response []byte
	if strings.Contains(strings.ToLower(requestContent), "ping") {
		response = []byte("+PONG\r\n")
	}

	_, err = conn.Write(response)
	if err != nil {
		fmt.Println("found an error trying to respond")
		return
	}
}
