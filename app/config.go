package main

type RedisConfig struct {
	isReplica bool
}

func (rs *RedisConfig) SetIsReplica(isReplica bool) {
	rs.isReplica = isReplica
}

func StartRedisConfig() *RedisConfig {
	return &RedisConfig{isReplica: false}
}

func (rs RedisConfig) GetRole() string {
	if rs.isReplica {
		return "slave"
	}
	return "master"
}
