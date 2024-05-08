package operations

type RedisOperation interface {
	HandleOperation() (string, error)
}
