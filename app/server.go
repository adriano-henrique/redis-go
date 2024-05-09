package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/codecrafters-io/redis-starter-go/app/utils"
)

func main() {
	redisStorage := utils.StartRedisStorage()
	redisConfig := utils.StartRedisConfig()

	var portFlag string
	flag.StringVar(&portFlag, "port", "6379", "The port in which you wish to bind the redis service to")

	replicaOf := flag.String("replicaof", "", "replicate to another redis server")
	flag.Parse()

	if *replicaOf != "" {
		redisConfig.SetIsReplica(true)
	}
	redisConfig.ConfigRedis()
	l, err := net.Listen("tcp", "0.0.0.0:"+portFlag)
	if err != nil {
		fmt.Println("Failed to bind to port " + portFlag)
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn, redisStorage, redisConfig)
	}
}

func handleConnection(conn net.Conn, storage *utils.RedisStorage, config *utils.RedisConfig) {
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
