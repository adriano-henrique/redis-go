package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var validOperations = []string{"ping", "echo"}

func ParseElements(elements []string) []string {
	var responses []string
	for i, element := range elements {
		lowerCaseValue := strings.ToLower(element)
		if len(strings.Split(element, " ")) == 1 && slices.Contains(validOperations, lowerCaseValue) {
			if lowerCaseValue == "ping" {
				responses = append(responses, "+PONG\r\n")
			} else if lowerCaseValue == "echo" {
				if i+1 >= len(elements) {
					fmt.Println("Should have peek element")
					os.Exit(1)
				}
				peekElement := elements[i+1]
				echoOutput := fmt.Sprintf("$%d\r\n%s\r\n", len(peekElement), peekElement)
				responses = append(responses, echoOutput)
			}
		}
	}
	return responses
}
