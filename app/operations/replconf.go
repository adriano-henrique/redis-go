package operations

import "github.com/codecrafters-io/redis-starter-go/app/utils"

type ReplConfOperation struct{}

func (rco ReplConfOperation) HandleOperation() (string, error) {
	replConfResponse := utils.NewRedisResponse(utils.Ok, utils.EmptyList())
	return replConfResponse.GetEncodedResponse()
}
