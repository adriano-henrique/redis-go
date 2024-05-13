package operations

import (
	"encoding/base64"
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

func (po PsyncOperation) HandleOperationMultipleResponses() ([]string, error) {
	var responses []string
	psyncResponse := utils.NewRedisResponse(utils.SimpleString, []string{po.buildResponseString()})
	psyncResponseString, err := psyncResponse.GetEncodedResponse()
	if err != nil {
		return []string{}, err
	}
	responses = append(responses, psyncResponseString)
	emptyFileStringRDB, err := po.buildEmptyFileString()
	if err != nil {
		return []string{}, err
	}
	rdbFileResponse := utils.NewRedisResponse(utils.FileResponse, []string{emptyFileStringRDB})
	rdbFileResponseString, err := rdbFileResponse.GetEncodedResponse()
	if err != nil {
		return []string{}, err
	}
	responses = append(responses, rdbFileResponseString)
	return responses, nil
}

func (po PsyncOperation) buildResponseString() string {
	return fmt.Sprintf("FULLRESYNC %s %d", po.config.MasterReplId, po.config.MasterReplOffset)
}

func (po PsyncOperation) buildEmptyFileString() (string, error) {
	decodedString, err := base64.StdEncoding.DecodeString(utils.RDBEmptyFileBase64())
	if err != nil {
		fmt.Println("Unable to decode string")
		return "", nil
	}
	return string(decodedString), nil
}
