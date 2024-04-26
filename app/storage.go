package main

import (
	"errors"
	"time"
)

type StorageObject struct {
	value  string
	expiry time.Time
}

func (object StorageObject) hasExpired() bool {
	if object.expiry.Compare(defaultTime) == 0 {
		return false
	}

	if time.Now().Compare(object.expiry) > 0 {
		return true
	}
	return false
}

type RedisStorage struct {
	data map[string]StorageObject
}

func StartRedisStorage() *RedisStorage {
	return &RedisStorage{data: make(map[string]StorageObject)}
}

func (rs RedisStorage) Get(key string) (StorageObject, error) {
	storageObject, ok := rs.data[key]
	if !ok {
		return StorageObject{}, errors.New("value not in map")
	}
	return storageObject, nil
}

func (rs *RedisStorage) Delete(key string) {
	delete(rs.data, key)
}

func (rs *RedisStorage) Set(key string, value StorageObject) error {
	rs.data[key] = value
	return nil
}
