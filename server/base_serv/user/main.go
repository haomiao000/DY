package main

import (
	"fmt"
	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	discovery "github.com/haomiao000/DY/comm/discovery"
	redis "github.com/haomiao000/DY/comm/redis"
	trace "github.com/haomiao000/DY/comm/trace"
	rpc_user "github.com/haomiao000/DY/internal/grpc_gen/rpc_user"
	interceptor "github.com/haomiao000/DY/internal/interceptor"
	api_server "github.com/haomiao000/DY/server/base_serv/user/api_server"
	configs "github.com/haomiao000/DY/server/base_serv/user/configs"
	dao "github.com/haomiao000/DY/server/base_serv/user/dao"
	initialize "github.com/haomiao000/DY/server/base_serv/user/initialize"
	grpc "google.golang.org/grpc"
)

func main() {
	db := initialize.InitDB()
	initialize.InitSnowFlake()
	discovery.Init()
	redis.Init()
	tracer, closer := trace.NewTracer("user")
	// All the bells and whistles:
	defer closer.Close()
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(
		grpcMiddleware.ChainUnaryServer(
			interceptor.ServerInterceptor(tracer),
		),
	))
	impl := &api_server.UserServiceImpl{
		MysqlManager: dao.NewMysqlManager(db),	
	}
	rpc_user.RegisterUserServiceImplServer(grpcServer, impl)
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", configs.GetPort()))
	if err != nil {
		fmt.Printf("Failed To Listen On Addr %v",err)
	}
	if err := discovery.Register(configs.GetServiceName(), configs.GetAddress()); err != nil {
		fmt.Printf("Failed To Init GRPC Server: %v", err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Printf("Failed To Serve GRPC Server: %v", err)
	}
}
