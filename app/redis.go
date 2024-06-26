package main

import (
	"fmt"
	"net"
	"os"
	"slices"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/app/connection"
	"github.com/codecrafters-io/redis-starter-go/app/operations"
	"github.com/codecrafters-io/redis-starter-go/app/utils"
)

var validOperations = []string{"ping", "echo", "get", "set", "info", "replconf", "psync"}

func AppendResponse(responses *[]string, operation operations.RedisOperation) {
	response, err := operation.HandleOperation()
	if err != nil {
		fmt.Println("Error during operation parsing: ", err.Error())
		os.Exit(1)
	}
	*responses = append(*responses, response)
}

func AppendResponses(responses *[]string, operation operations.RedisOperation) {
	additionalResponses, err := operation.HandleOperationMultipleResponses()
	if err != nil {
		fmt.Println("Error during operation parsing: ", err.Error())
		os.Exit(1)
	}
	*responses = append(*responses, additionalResponses...)
}

func ParseElements(elements []string, storage *utils.RedisStorage, config *utils.RedisConfig, conn net.Conn) []string {
	var responses []string
	for i, element := range elements {
		lowerCaseValue := strings.ToLower(element)
		if len(strings.Split(element, " ")) == 1 && slices.Contains(validOperations, lowerCaseValue) {
			switch command := lowerCaseValue; command {
			case "ping":
				pingOperation := operations.PingOperation{}
				AppendResponse(&responses, pingOperation)
			case "echo":
				echoOperation := operations.NewEchoOperation(i, elements)
				AppendResponse(&responses, echoOperation)
			case "get":
				getOperation := operations.NewGetOperation(i, elements, storage)
				AppendResponse(&responses, getOperation)
				propagateCommand(config, elements)
			case "set":
				setOperation := operations.NewSetOperation(i, elements, storage)
				AppendResponse(&responses, setOperation)
				propagateCommand(config, elements)
			case "info":
				infoOperation := operations.NewInfoOperation(config)
				AppendResponse(&responses, infoOperation)
			case "replconf":
				replConfOperation := operations.ReplConfOperation{}
				AppendResponse(&responses, replConfOperation)
			case "psync":
				psyncOperation := operations.NewPsyncOperation(config)
				AppendResponses(&responses, psyncOperation)
				config.AddReplica(conn)
			}
		}
	}
	return responses
}

func propagateCommand(config *utils.RedisConfig, elements []string) {
	for _, conn := range config.ReplicaConnections {
		conn.Write(connection.WriteBulkString(elements))
	}
}
