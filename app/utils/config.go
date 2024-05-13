package utils

import "fmt"

type RedisConfig struct {
	isReplica        bool
	MasterHost       string
	MasterReplId     string
	MasterReplOffset int
}

func (rs *RedisConfig) SetIsReplica(isReplica bool) {
	rs.isReplica = isReplica
}

func (rs *RedisConfig) SetMasterHostAddress(hostAddress string) {
	rs.MasterHost = hostAddress
}

func StartRedisConfig() *RedisConfig {
	return &RedisConfig{
		isReplica:        false,
		MasterHost:       "",
		MasterReplId:     "",
		MasterReplOffset: 0,
	}
}

func (rs *RedisConfig) ConfigRedis() {
	rs.MasterReplId = "8371b4fb1155b71f4a04d3e1bc3e18c4a990aeeb"
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
	return fmt.Sprintf("master_replid:%s", rs.MasterReplId)
}

func (rs RedisConfig) GetMasterReplOffsetString() string {
	return fmt.Sprintf("master_repl_offset:%d", rs.MasterReplOffset)
}
