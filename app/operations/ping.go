package operations

type PingOperation struct{}

func (p PingOperation) HandleOperation() (string, error) {
	pongResponse := &RedisResponse{
		responseType: Ok,
		elements:     EmptyList(),
	}

	return pongResponse.GetEncodedResponse()
}
