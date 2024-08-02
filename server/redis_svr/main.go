package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/haomiao000/DY/server/redis_svr/api"
	"github.com/haomiao000/DY/server/redis_svr/internal"
	pb "github.com/haomiao000/DY/server/redis_svr/pb/redis_svr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := internal.Init(); err != nil {
		fmt.Println("init error: err", err)
		return
	}
	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		fmt.Println(err)
		return
	}
	s := grpc.NewServer()
	impl := &api.RedisSvrImpl{}
	pb.RegisterRedisSvrServer(s, impl)
	go func() {
		time.Sleep(time.Second * 2)
		con, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		redisCli := pb.NewRedisSvrClient(con)
		fmt.Println(redisCli.Get(context.Background(), &pb.GetReq{Key: "1"}))
	}()
	if err = s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
