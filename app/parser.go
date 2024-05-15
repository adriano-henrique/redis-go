package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseRequest(requestContent string) []string {
	numElements, err := strconv.Atoi(string(requestContent[1]))
	if err != nil {
		fmt.Println("Unable to parse number of elements - should be number. Error: ", err.Error())
		os.Exit(1)
	}
	requestElementsList := make([]string, numElements)
	request := getRequestElementsList(requestContent)

	currDataIndex := 1
	for i := 0; i < numElements; i++ {
		currentElement := request[currDataIndex]
		if currDataIndex+1 >= len(request) {
			fmt.Println("Invalid format, should have a data after the number of elements")
			os.Exit(1)
		}
		peekElement := request[currDataIndex+1]
		var elementSize int
		if currentElement[0] == '$' {
			elementSize, err = strconv.Atoi(string(currentElement[1:]))
			newElement := make([]byte, 0, elementSize)
			if err != nil {
				fmt.Println("Unable to get size of data. Error: ", err.Error())
				os.Exit(1)
			}
			inputBytes := []byte(peekElement)

			if len(inputBytes) != elementSize {
				panic("length of input string doesn't match the expected number of bytes")
			}
			if string(inputBytes) != "" {
				newElement = append(newElement, inputBytes...)
			}
			if string(newElement) != "" {
				requestElementsList = append(requestElementsList, string(newElement))
			}
		}
		currDataIndex += 2
	}

	return removeWhitespace(requestElementsList)
}

func removeWhitespace(stringList []string) []string {
	var ans []string
	for _, i := range stringList {
		if i != "" {
			ans = append(ans, i)
		}
	}
	return ans
}

func getRequestElementsList(requestContent string) []string {
	if requestContent[0] != byte('*') {
		fmt.Println("Invalid request to Redis - should start with *")
		os.Exit(1)
	}
	requestEnd := strings.LastIndex(requestContent, "\r\n")
	var request string
	if requestEnd != -1 {
		request = requestContent[:requestEnd+2]
	} else {
		request = requestContent
	}

	return strings.Split(request, "\r\n")
}
