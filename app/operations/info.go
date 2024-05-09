package operations

import (
	"strings"

	"github.com/codecrafters-io/redis-starter-go/app/utils"
)

type InfoOperation struct {
	config *utils.RedisConfig
}

func NewInfoOperation(config *utils.RedisConfig) *InfoOperation {
	return &InfoOperation{
		config: config,
	}
}

func (io InfoOperation) HandleOperation() (string, error) {
	infoResponse := utils.NewRedisResponse(utils.SingleElement, []string{io.buildResponseString()})
	return infoResponse.GetEncodedResponse()
}

func (io InfoOperation) buildResponseElements() []string {
	return []string{
		io.config.GetRoleInfoString(),
		io.config.GetMasterReplIdString(),
		io.config.GetMasterReplOffsetString(),
	}
}

func (io InfoOperation) buildResponseString() string {
	return strings.Join(io.buildResponseElements(), "\n")
}
