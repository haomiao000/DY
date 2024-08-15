package main

import (
	"fmt"
	"net"
	rpc_user "github.com/haomiao000/DY/internal/grpc_gen/rpc_user"
	api_server "github.com/haomiao000/DY/server/base_serv/user/api_server"
	configs "github.com/haomiao000/DY/server/base_serv/user/configs"
	dao "github.com/haomiao000/DY/server/base_serv/user/dao"
	initialize "github.com/haomiao000/DY/server/base_serv/user/initialize"
	grpc "google.golang.org/grpc"
	discovery "github.com/haomiao000/DY/comm/discovery"
	redis "github.com/haomiao000/DY/comm/redis"
)
const (
    // 服务名 服务唯一标示，服务指标聚合过滤依据。
    grpcServerName = "Register"
)

func main() {
	db := initialize.InitDB()
	discovery.Init()
	redis.Init()
	grpcServer := grpc.NewServer()
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
