package utils

import "fmt"

type RedisConfig struct {
	isReplica bool
}

func (rs *RedisConfig) SetIsReplica(isReplica bool) {
	rs.isReplica = isReplica
}

func StartRedisConfig() *RedisConfig {
	return &RedisConfig{isReplica: false}
}

func (rs RedisConfig) getRole() string {
	if rs.isReplica {
		return "slave"
	}
	return "master"
}

func (rs RedisConfig) GetRoleInfoString() string {
	return fmt.Sprintf("role:%s", rs.getRole())
}
