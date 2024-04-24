package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

var validOperations = []string{"ping", "echo", "get", "set"}

func ParseElements(elements []string, storage *RedisStorage) []string {
	var responses []string
	for i, element := range elements {
		lowerCaseValue := strings.ToLower(element)
		if len(strings.Split(element, " ")) == 1 && slices.Contains(validOperations, lowerCaseValue) {
			switch command := lowerCaseValue; command {
			case "ping":
				response := handlePing()
				responses = append(responses, response)
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
	value, err := storage.Get(keyElement)
	if err != nil {
		return "$-1\r\n", nil
	}
	return fmt.Sprintf("$%d\r\n%s\r\n", len(value), value), nil
}

func handleSet(index int, elements []string, storage *RedisStorage) (string, error) {
	if index+1 >= len(elements) || index+2 >= len(elements) {
		return "", errors.New("expected to have key and value")
	}
	keyElement := elements[index+1]
	valueElement := elements[index+2]
	storage.Set(keyElement, valueElement)
	return "+OK\r\n", nil
}
