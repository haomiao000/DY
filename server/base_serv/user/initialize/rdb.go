package initialize

import (
	configs "github.com/haomiao000/DY/server/base_serv/user/configs"
	grpc "google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"
	redis "github.com/haomiao000/DY/server/redis_svr/pb/redis_svr"
)

func InitRedis() redis.RedisSvrClient {
	con, err := grpc.Dial(configs.RedisIP+configs.RedisPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return redis.NewRedisSvrClient(con)
}

// func InitRedis() *redigo.Conn {
// 	c, err := redigo.Dial("tcp", configs.RedisIP+configs.RedisPort)
// 	if err != nil {
// 		return nil
// 	}
// 	return &c
// }