package main

import (
	"fmt"
	"net"

	rpc_relation "github.com/haomiao000/DY/internal/grpc_gen/rpc_relation"
	api_client "github.com/haomiao000/DY/server/base_serv/relation/api_client"
	api_server "github.com/haomiao000/DY/server/base_serv/relation/api_server"
	configs "github.com/haomiao000/DY/server/base_serv/relation/configs"
	dao "github.com/haomiao000/DY/server/base_serv/relation/dao"
	initialize "github.com/haomiao000/DY/server/base_serv/relation/initialize"
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
	tracer, closer := trace.NewTracer("relation")
	defer closer.Close()
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(
		grpcMiddleware.ChainUnaryServer(
			interceptor.ServerInterceptor(tracer),
		),
	))
	impl := &api_server.RelationServiceImpl{
		MysqlManager: dao.NewMysqlManager(db),
		UserManager:  api_client.NewUserClient(userServ),
	}
	rpc_relation.RegisterRelationServiceImplServer(grpcServer, impl)
	if err := discovery.Register(configs.GetServiceName(), configs.GetAddress()); err != nil {
		fmt.Printf("Failed To Init GRPC Server: %v", err)
	}
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", configs.GetPort()))
	if err != nil {
		fmt.Printf("Failed To Listen On Addr: %v", err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Printf("Failed To Serve GRPC Server: %v", err)
	}
}
