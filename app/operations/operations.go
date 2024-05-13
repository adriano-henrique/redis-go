package operations

type RedisOperation interface {
	HandleOperation() (string, error)
	HandleOperationMultipleResponses() ([]string, error)
}
