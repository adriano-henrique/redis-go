package operations

import "github.com/codecrafters-io/redis-starter-go/app/utils"

type PingOperation struct{}

func (p PingOperation) HandleOperation() (string, error) {
	pongResponse := utils.NewRedisResponse(utils.SimpleString, []string{"PONG"})
	return pongResponse.GetEncodedResponse()
}

func (p PingOperation) HandleOperationMultipleResponses() ([]string, error) {
	return []string{}, nil
}
