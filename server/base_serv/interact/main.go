package main

import (
	"fmt"
	"net"

	rpc_interact "github.com/haomiao000/DY/internal/grpc_gen/rpc_interact"
	api_client "github.com/haomiao000/DY/server/base_serv/interact/api_client"
	api_server "github.com/haomiao000/DY/server/base_serv/interact/api_server"
	configs "github.com/haomiao000/DY/server/base_serv/interact/configs"
	dao "github.com/haomiao000/DY/server/base_serv/interact/dao"
	initialize "github.com/haomiao000/DY/server/base_serv/interact/initialize"
	grpc "google.golang.org/grpc"
	discovery "github.com/haomiao000/DY/comm/discovery"
	redis "github.com/haomiao000/DY/comm/redis"
	trace "github.com/haomiao000/DY/comm/trace"
	interceptor "github.com/haomiao000/DY/internal/interceptor"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
)

func main() {
	db := initialize.InitDB()
	discovery.Init()
	redis.Init()
	userServ := initialize.InitUser()
	videoServ := initialize.InitVideo()
	tracer, closer := trace.NewTracer("interact")
	defer closer.Close()
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(
		grpcMiddleware.ChainUnaryServer(
			interceptor.ServerInterceptor(tracer),
		),
	))
	impl := &api_server.InteractServiceImpl{
		FavoriteMysqlManager: dao.NewMysqlManager(db),
		CommentMysqlManager:  dao.NewMysqlManager(db),
		UserManager:          api_client.NewUserClient(userServ),
		VideoManager:         api_client.NewVideoClient(videoServ),
	}
	rpc_interact.RegisterInteractServiceImplServer(grpcServer, impl)
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
