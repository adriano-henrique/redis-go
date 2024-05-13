package operations

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/app/utils"
)

type PsyncOperation struct {
	config *utils.RedisConfig
}

func NewPsyncOperation(config *utils.RedisConfig) *PsyncOperation {
	return &PsyncOperation{
		config: config,
	}
}

func (po PsyncOperation) HandleOperation() (string, error) {
	psyncResponse := utils.NewRedisResponse(utils.SimpleString, []string{po.buildResponseString()})
	return psyncResponse.GetEncodedResponse()
}

func (po PsyncOperation) buildResponseString() string {
	return fmt.Sprintf("FULLRESYNC %s %d", po.config.MasterReplId, po.config.MasterReplOffset)
}
