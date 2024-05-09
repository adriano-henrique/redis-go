package operations

import (
	"errors"

	"github.com/codecrafters-io/redis-starter-go/app/utils"
)

type EchoOperation struct {
	index    int
	elements []string
}

func NewEchoOperation(index int, elements []string) *EchoOperation {
	return &EchoOperation{
		index:    index,
		elements: elements,
	}
}

func (eo EchoOperation) HandleOperation() (string, error) {
	if eo.index+1 >= len(eo.elements) {
		return "", errors.New("should have a peek element (string to be echoed)")
	}
	peekElement := eo.elements[eo.index+1]
	echoResponse := utils.NewRedisResponse(utils.BulkString, []string{peekElement})
	return echoResponse.GetEncodedResponse()
}
