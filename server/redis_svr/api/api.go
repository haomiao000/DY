package api

import pb "github.com/haomiao000/DY/server/redis_svr/pb/redis_svr"
// 实现接口结构体
type RedisSvrImpl struct{
	pb.UnimplementedRedisSvrServer
}
