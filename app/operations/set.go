package operations

import (
	"errors"
	"strconv"
	"time"

	"github.com/codecrafters-io/redis-starter-go/app/utils"
)

type SetOperation struct {
	index    int
	elements []string
	storage  *utils.RedisStorage
}

func NewSetOperation(index int, elements []string, storage *utils.RedisStorage) *SetOperation {
	return &SetOperation{
		index:    index,
		elements: elements,
		storage:  storage,
	}
}

func (so SetOperation) HandleOperation() (string, error) {
	if so.index+1 >= len(so.elements) || so.index+2 >= len(so.elements) {
		return "", errors.New("expected to have key and value")
	}
	keyElement := so.elements[so.index+1]
	valueElement := so.elements[so.index+2]
	if so.index+3 < len(so.elements) && so.elements[so.index+3] == "px" {
		if so.index+4 >= len(so.elements) {
			return "", errors.New("should have timeout value if has px command")
		}

		timeout, err := strconv.Atoi(so.elements[so.index+4])
		if err != nil {
			return "", err
		}
		milisecondsAdded := time.Duration(timeout) * time.Millisecond
		object := utils.NewStorageObject(valueElement, time.Now().Add(milisecondsAdded))
		so.storage.Set(keyElement, object)
		return utils.OkResponse().GetEncodedResponse()
	}
	storedObject := utils.NewStorageObject(valueElement, utils.DefaultTime())
	so.storage.Set(keyElement, storedObject)
	return utils.OkResponse().GetEncodedResponse()
}

func (so SetOperation) HandleOperationMultipleResponses() ([]string, error) {
	return []string{}, nil
}
