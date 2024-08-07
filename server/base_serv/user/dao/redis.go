package dao

import (
	redis "github.com/haomiao000/DY/server/redis_svr/pb/redis_svr"
	
	"context"
)

type RedisManager struct {
	redisClient redis.RedisSvrClient
}

func (r *RedisManager) SetUserLoginInfo(ctx context.Context , key string , val string) error {
	_ , err := r.redisClient.SetWithExpire(ctx, &redis.SetWithExpireReq{
		Key: key,
		Val: val,
		Expire: 0,
	})
	if err != nil {
		return err
	}
	return nil
}



func NewRedisSvrClient(rdb redis.RedisSvrClient) *RedisManager {
	return &RedisManager{redisClient: rdb}
}