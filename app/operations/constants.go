package operations

func EmptyList() []string {
	var emptyList []string
	return emptyList
}

func ErrorResponse() *RedisResponse {
	return &RedisResponse{
		responseType: Invalid,
		elements:     EmptyList(),
	}
}

func OkResponse() *RedisResponse {
	return &RedisResponse{
		responseType: Ok,
		elements:     EmptyList(),
	}
}
