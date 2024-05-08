package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/app/operations"
	"github.com/codecrafters-io/redis-starter-go/app/utils"
)

var validOperations = []string{"ping", "echo", "get", "set", "info"}

func AppendResponse(responses *[]string, operation operations.RedisOperation) {
	response, err := operation.HandleOperation()
	if err != nil {
		fmt.Println("Error during operation partsing: ", err.Error())
		os.Exit(1)
	}
	*responses = append(*responses, response)
}

func ParseElements(elements []string, storage *utils.RedisStorage, config *RedisConfig) []string {
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
			case "set":
				setOperation := operations.NewSetOperation(i, elements, storage)
				AppendResponse(&responses, setOperation)
			case "info":
				response, err := handleInfo(config)
				if err != nil {
					fmt.Println("Error during operation partsing: ", err.Error())
					os.Exit(1)
				}
				responses = append(responses, response)
			}
		}
	}
	return responses
}

func handleInfo(config *RedisConfig) (string, error) {

	return "$11\r\nrole:master\r\n", nil
}
