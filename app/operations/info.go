package operations

import "github.com/codecrafters-io/redis-starter-go/app/utils"

type InfoOperation struct{}

func NewInfoOperation() *InfoOperation {
	return &InfoOperation{}
}

func (io InfoOperation) HandleOperation() (string, error) {
	infoResponse := utils.NewRedisResponse(utils.SingleElement, []string{"role:master"})
	return infoResponse.GetEncodedResponse()
}
