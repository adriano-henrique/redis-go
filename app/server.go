package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	redisStorage := StartRedisStorage()
	redisConfig := StartRedisConfig()

	var portFlag string
	flag.StringVar(&portFlag, "port", "6379", "The port in which you wish to bind the redis service to")

	l, err := net.Listen("tcp", "0.0.0.0:"+portFlag)
	if err != nil {
		fmt.Println("Failed to bind to port " + portFlag)
		os.Exit(1)
	}
	defer l.Close()

	var replicaOf string
	flag.StringVar(&replicaOf, "replicaof", "", "The redis instance you wish to make this a slave to")
	if replicaOf != "" {
		redisConfig.SetIsReplica(true)
	}

	flag.Parse()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn, redisStorage, redisConfig)
	}
}

func handleConnection(conn net.Conn, storage *RedisStorage, config *RedisConfig) {
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
		responses := ParseElements(requestElements, storage, config)
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
