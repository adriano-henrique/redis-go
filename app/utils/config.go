package utils

import "fmt"

type RedisConfig struct {
	isReplica        bool
	masterReplId     string
	masterReplOffset int
}

func (rs *RedisConfig) SetIsReplica(isReplica bool) {
	rs.isReplica = isReplica
}

func StartRedisConfig() *RedisConfig {
	return &RedisConfig{
		isReplica:        false,
		masterReplId:     "",
		masterReplOffset: 0,
	}
}

func (rs *RedisConfig) ConfigRedis() {
	rs.masterReplId = "8371b4fb1155b71f4a04d3e1bc3e18c4a990aeeb"
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

func (rs RedisConfig) GetMasterReplIdString() string {
	return fmt.Sprintf("master_replid:%s", rs.masterReplId)
}

func (rs RedisConfig) GetMasterReplOffsetString() string {
	return fmt.Sprintf("master_repl_offset:%d", rs.masterReplOffset)
}
