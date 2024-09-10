package main

import (
	"fmt"
	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	discovery "github.com/haomiao000/DY/comm/discovery"
	redis "github.com/haomiao000/DY/comm/redis"
	trace "github.com/haomiao000/DY/comm/trace"
	rpc_video "github.com/haomiao000/DY/internal/grpc_gen/rpc_video"
	interceptor "github.com/haomiao000/DY/internal/interceptor"
	"github.com/haomiao000/DY/server/base_serv/video/api_client"
	api_server "github.com/haomiao000/DY/server/base_serv/video/api_server"
	configs "github.com/haomiao000/DY/server/base_serv/video/configs"
	dao "github.com/haomiao000/DY/server/base_serv/video/dao"
	initialize "github.com/haomiao000/DY/server/base_serv/video/initialize"
	grpc "google.golang.org/grpc"
)

func main() {
	db := initialize.InitDB()
	if err := discovery.Init(); err != nil {
		panic(err)
	}
	if err := redis.Init(); err != nil {
		panic(err)
	}
	if err := api_client.Init(); err != nil {
		panic(err)
	}
	tracer, closer := trace.NewTracer("video")
	defer closer.Close()
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(
		grpcMiddleware.ChainUnaryServer(
			interceptor.ServerInterceptor(tracer),
		),
	))
	impl := &api_server.VideoServiceImpl{
		MysqlManager: dao.NewMysqlManager(db),
	}
	rpc_video.RegisterVideoServiceImplServer(grpcServer, impl)
	if err := discovery.Register(configs.GetServiceName(), configs.GetAddress()); err != nil {
		fmt.Printf("Failed To Init GRPC Server: %v", err)
	}
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", configs.GetPort()))
	if err != nil {
		fmt.Printf("Failed To Listen On Addr %v", err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Printf("Failed To Serve GRPC Server: %v", err)
	}
}
