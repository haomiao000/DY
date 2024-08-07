package dao

import (
	redis "github.com/haomiao000/DY/server/redis_svr/pb/redis_svr"
)

type RedisManager struct {
	redisClient redis.RedisSvrClient
}



func NewRedisSvrClient(rdb redis.RedisSvrClient) *RedisManager {
	return &RedisManager{redisClient: rdb}
}