package main

import (
	"fmt"
	"net"

	"github.com/haomiao000/DY/server/redis_svr/api"
	"github.com/haomiao000/DY/server/redis_svr/internal"
	pb "github.com/haomiao000/DY/server/redis_svr/pb/redis_svr"
	"google.golang.org/grpc"
)

func main() {
	if err := internal.Init(); err != nil {
		fmt.Println("init error: err", err)
		return
	}
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		fmt.Println(err)
		return
	}
	s := grpc.NewServer()
	impl := &api.RedisSvrImpl{}
	pb.RegisterRedisSvrServer(s, impl)
	if err = s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
