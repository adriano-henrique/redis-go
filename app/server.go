package main

import (
	"fmt"
	"io"
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
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("Got an error when trying to read: ", err.Error())
			os.Exit(1)
		}

		requestContent := string(readBuffer)
		requestElements := ParseRequest(requestContent)
		responses := ParseElements(requestElements)
		fmt.Println(responses)

		for _, response := range responses {
			res := []byte(response)
			n, err := conn.Write(res)
			if err != nil || n != len(response) {
				fmt.Println("found an error trying to respond")
				return
			}
		}
	}
}
