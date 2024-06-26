package utils

import (
	"errors"
	"fmt"
	"strings"
)

type ResponseType int64

const (
	Ok ResponseType = iota
	Invalid
	BulkString
	ArrayResponse
	SimpleString
	FileResponse
)

type RedisResponse struct {
	responseType ResponseType
	elements     []string
}

func NewRedisResponse(responseType ResponseType, elements []string) *RedisResponse {
	return &RedisResponse{
		responseType: responseType,
		elements:     elements,
	}
}

func (rs *RedisResponse) GetEncodedResponse() (string, error) {
	switch rs.responseType {
	case Ok:
		return encodeSimpleString("OK"), nil
	case Invalid:
		return "$-1\r\n", nil
	case SimpleString:
		if len(rs.elements) == 0 {
			return "", errors.New("elements to build the response should be passed")
		}
		element := rs.elements[0]
		return encodeSimpleString(element), nil
	case BulkString:
		if len(rs.elements) == 0 {
			return "", errors.New("elements to build the response should be passed")
		}
		element := rs.elements[0]
		return encodeRedisBulkString(element), nil
	case FileResponse:
		if len(rs.elements) == 0 {
			return "", errors.New("elements to build the response should be passed")
		}
		element := rs.elements[0]
		return encodeFileString(element), nil
	case ArrayResponse:
		if len(rs.elements) <= 1 {
			return "", errors.New("incorrect value passed, expected more than 1 element")
		}
		numElements := len(rs.elements)
		var responseBuilder strings.Builder
		prefix := fmt.Sprintf("*%d\r\n", numElements)
		responseBuilder.WriteString(prefix)
		var responseBodyBuilder strings.Builder
		for _, element := range rs.elements {
			responseBodyBuilder.WriteString(encodeRedisBulkString(element))
		}
		responseBody := responseBodyBuilder.String()
		responseBuilder.WriteString(responseBody)
		return responseBuilder.String(), nil
	}

	return "", errors.New("invalid reponse type")
}

func encodeFileString(value string) string {
	return fmt.Sprintf("$%d\r\n%s", len(value), value)
}

func encodeRedisBulkString(value string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(value), value)
}

func encodeSimpleString(value string) string {
	return fmt.Sprintf("+%s\r\n", value)
}
