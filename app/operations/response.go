package operations

import (
	"errors"
	"fmt"
)

type ResponseType int64

const (
	Ok ResponseType = iota
	Pong
	Invalid
	SingleElement
)

type RedisResponse struct {
	responseType ResponseType
	elements     []string
}

func (rs *RedisResponse) GetEncodedResponse() (string, error) {
	switch rs.responseType {
	case Ok:
		return "+OK\r\n", nil
	case Invalid:
		return "$-1\r\n", nil
	case Pong:
		return "+PONG\r\n", nil
	case SingleElement:
		if len(rs.elements) == 0 {
			return "", errors.New("elements to build the response should be passed")
		}
		element := rs.elements[0]
		return fmt.Sprintf("$%d\r\n%s\r\n", len(element), element), nil
	}

	return "", errors.New("invalid reponse type")
}
