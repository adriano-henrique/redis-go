package utils

import (
	"errors"
	"time"
)

type StorageObject struct {
	value  string
	expiry time.Time
}

var defaultTime = time.Unix(0, 0)

func (object StorageObject) HasExpired() bool {
	if object.expiry.Compare(defaultTime) == 0 {
		return false
	}

	if time.Now().Compare(object.expiry) > 0 {
		return true
	}
	return false
}

func (object StorageObject) Value() string {
	return object.value
}

func (object StorageObject) Expiry() time.Time {
	return object.expiry
}

func NewStorageObject(value string, expiry time.Time) *StorageObject {
	return &StorageObject{
		value:  value,
		expiry: expiry,
	}
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

func (rs *RedisStorage) Set(key string, value *StorageObject) error {
	rs.data[key] = *value
	return nil
}
