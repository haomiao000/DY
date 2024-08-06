package main

import (
	"fmt"
	"net"

	"github.com/haomiao000/DY/comm/discovery"
	"github.com/haomiao000/DY/server/redis_svr/api"
	"github.com/haomiao000/DY/server/redis_svr/config"
	"github.com/haomiao000/DY/server/redis_svr/internal"
	pb "github.com/haomiao000/DY/server/redis_svr/pb/redis_svr"
	"google.golang.org/grpc"
)

func main() {
	if err := internal.Init(); err != nil {
		fmt.Printf("init error: %v", err)
		return
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.GetPort()))
	if err != nil {
		fmt.Println(err)
		return
	}
	s := grpc.NewServer()
	impl := &api.RedisSvrImpl{}
	pb.RegisterRedisSvrServer(s, impl)
	discovery.Register(config.GetServiceName(), config.GetAddress())
	if err = s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
