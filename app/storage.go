package main

import "errors"

type RedisStorage struct {
	data map[string]string
}

func StartRedisStorage() *RedisStorage {
	return &RedisStorage{data: make(map[string]string)}
}

func (rs RedisStorage) Get(key string) (string, error) {
	val, ok := rs.data[key]
	if !ok {
		return "", errors.New("value not in map")
	}
	return val, nil
}

func (rs *RedisStorage) Set(key string, value string) error {
	rs.data[key] = value
	return nil
}
