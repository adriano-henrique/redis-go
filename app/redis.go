package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/codecrafters-io/redis-starter-go/app/operations"
)

var validOperations = []string{"ping", "echo", "get", "set", "info"}
var defaultTime = time.Unix(0, 0)

func AppendResponse(responses *[]string, operation operations.RedisOperation) {
	response, err := operation.HandleOperation()
	if err != nil {
		fmt.Println("Error during operation partsing: ", err.Error())
		os.Exit(1)
	}
	*responses = append(*responses, response)
}

func ParseElements(elements []string, storage *RedisStorage, config *RedisConfig) []string {
	var responses []string
	for i, element := range elements {
		lowerCaseValue := strings.ToLower(element)
		if len(strings.Split(element, " ")) == 1 && slices.Contains(validOperations, lowerCaseValue) {
			switch command := lowerCaseValue; command {
			case "ping":
				pingOperation := operations.PingOperation{}
				AppendResponse(&responses, pingOperation)
			case "echo":
				response, err := handleEcho(i, elements)
				if err != nil {
					fmt.Println("Error during operation parsing: ", err.Error())
					os.Exit(1)
				}
				responses = append(responses, response)
			case "get":
				response, err := handleGet(i, elements, storage)
				if err != nil {
					fmt.Println("Error during operation parsing: ", err.Error())
					os.Exit(1)
				}
				responses = append(responses, response)
			case "set":
				response, err := handleSet(i, elements, storage)
				if err != nil {
					fmt.Println("Error during operation parsing: ", err.Error())
					os.Exit(1)
				}
				responses = append(responses, response)
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

func handlePing() string {
	return "+PONG\r\n"
}

func handleEcho(index int, elements []string) (string, error) {
	if index+1 >= len(elements) {
		return "", errors.New("should have a peek element (string to be echoed)")
	}
	peekElement := elements[index+1]
	return fmt.Sprintf("$%d\r\n%s\r\n", len(peekElement), peekElement), nil
}

func handleGet(index int, elements []string, storage *RedisStorage) (string, error) {
	if index+1 >= len(elements) {
		return "", errors.New("expected to have peek element (key to get from storage)")
	}
	keyElement := elements[index+1]
	object, err := storage.Get(keyElement)
	fmt.Println("Got: ", object.value)
	fmt.Println("Got: ", object.expiry.String())
	if err != nil {
		return "$-1\r\n", nil
	}

	if object.hasExpired() {
		storage.Delete(keyElement)
		return "$-1\r\n", nil
	}

	value := object.value
	return fmt.Sprintf("$%d\r\n%s\r\n", len(value), value), nil
}

func handleSet(index int, elements []string, storage *RedisStorage) (string, error) {
	if index+1 >= len(elements) || index+2 >= len(elements) {
		return "", errors.New("expected to have key and value")
	}
	keyElement := elements[index+1]
	valueElement := elements[index+2]
	if index+3 < len(elements) && elements[index+3] == "px" {
		if index+4 >= len(elements) {
			return "", errors.New("should have timeout value if has px command")
		}

		timeout, err := strconv.Atoi(elements[index+4])
		if err != nil {
			return "", err
		}
		milisecondsAdded := time.Duration(timeout) * time.Millisecond
		object := StorageObject{
			value:  valueElement,
			expiry: time.Now().Add(milisecondsAdded),
		}
		storage.Set(keyElement, object)
		return "+OK\r\n", nil
	}
	storedObject := StorageObject{
		value:  valueElement,
		expiry: defaultTime,
	}
	storage.Set(keyElement, storedObject)
	return "+OK\r\n", nil
}

func handleInfo(config *RedisConfig) (string, error) {

	return "$11\r\nrole:master\r\n", nil
}
