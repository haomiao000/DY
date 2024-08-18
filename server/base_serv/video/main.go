package main

import (
	rpc_video "github.com/haomiao000/DY/internal/grpc_gen/rpc_video"
	api_client "github.com/haomiao000/DY/server/base_serv/video/api_client"
	api_server "github.com/haomiao000/DY/server/base_serv/video/api_server"
	configs "github.com/haomiao000/DY/server/base_serv/video/configs"
	dao "github.com/haomiao000/DY/server/base_serv/video/dao"
	initialize "github.com/haomiao000/DY/server/base_serv/video/initialize"
	discovery "github.com/haomiao000/DY/comm/discovery"
	grpc "google.golang.org/grpc"
	redis "github.com/haomiao000/DY/comm/redis"
	"fmt"
	"net"
	trace "github.com/haomiao000/DY/comm/trace"
	interceptor "github.com/haomiao000/DY/internal/interceptor"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
)

func main() {
	db := initialize.InitDB()
	discovery.Init()
	redis.Init()
	tracer, closer := trace.NewTracer("video")
	defer closer.Close()
	userServ := initialize.InitUser()
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(
		grpcMiddleware.ChainUnaryServer(
			interceptor.ServerInterceptor(tracer),
		),
	))
	impl := &api_server.VideoServiceImpl{
		MysqlManager: dao.NewMysqlManager(db),
		UserManager:  api_client.NewUserClient(userServ),
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
