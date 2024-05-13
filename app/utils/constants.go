package utils

import "time"

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

func DefaultTime() time.Time {
	return time.Unix(0, 0)
}

func RDBEmptyFileBase64() string {
	return "UkVESVMwMDEx+glyZWRpcy12ZXIFNy4yLjD6CnJlZGlzLWJpdHPAQPoFY3RpbWXCbQi8ZfoIdXNlZC1tZW3CsMQQAPoIYW9mLWJhc2XAAP/wbjv+wP9aog=="
}
