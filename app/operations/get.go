package operations

import (
	"errors"

	"github.com/codecrafters-io/redis-starter-go/app/utils"
)

type GetOperation struct {
	index    int
	elements []string
	storage  *utils.RedisStorage
}

func NewGetOperation(index int, elements []string, storage *utils.RedisStorage) *GetOperation {
	return &GetOperation{
		index:    index,
		elements: elements,
		storage:  storage,
	}
}

func (gop GetOperation) HandleOperation() (string, error) {
	if gop.index+1 >= len(gop.elements) {
		return "", errors.New("expected to have peek element (key to get from storage)")
	}
	keyElement := gop.elements[gop.index+1]
	object, err := gop.storage.Get(keyElement)
	if err != nil {
		return "$-1\r\n", nil
	}

	if object.HasExpired() {
		gop.storage.Delete(keyElement)
		return "$-1\r\n", nil
	}

	getResponse := utils.NewRedisResponse(utils.BulkString, []string{object.Value()})
	return getResponse.GetEncodedResponse()
}

func (gop GetOperation) HandleOperationMultipleResponses() ([]string, error) {
	return []string{}, nil
}
