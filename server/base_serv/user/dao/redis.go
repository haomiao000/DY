package dao

import (
	redigo "github.com/gomodule/redigo/redis"
)

type RedisManager struct {
	redisClient *redigo.Conn
}



func NewRedisSvrClient(rdb *redigo.Conn) *RedisManager {
	return &RedisManager{redisClient: rdb}
}