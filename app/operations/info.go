package operations

import "github.com/codecrafters-io/redis-starter-go/app/utils"

type InfoOperation struct {
	config *utils.RedisConfig
}

func NewInfoOperation(config *utils.RedisConfig) *InfoOperation {
	return &InfoOperation{
		config: config,
	}
}

func (io InfoOperation) HandleOperation() (string, error) {
	infoResponse := utils.NewRedisResponse(utils.SingleElement, []string{io.config.GetRoleInfoString()})
	return infoResponse.GetEncodedResponse()
}
