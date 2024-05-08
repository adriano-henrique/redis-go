package operations

import "github.com/codecrafters-io/redis-starter-go/app/utils"

type PingOperation struct{}

func (p PingOperation) HandleOperation() (string, error) {
	pongResponse := utils.NewRedisResponse(utils.Pong, utils.EmptyList())
	return pongResponse.GetEncodedResponse()
}
