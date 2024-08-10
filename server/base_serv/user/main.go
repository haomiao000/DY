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
)

func main() {
	db := initialize.InitDB()
	grpcServer := grpc.NewServer()
	impl := &api_server.UserServiceImpl{
		MysqlManager: dao.NewMysqlManager(db),	
	}
	rpc_user.RegisterUserServiceImplServer(grpcServer, impl)

	listener, err := net.Listen("tcp", configs.UserServerAddr)
	if err != nil {
		fmt.Printf("Failed To Listen On Addr %s: %v", configs.UserServerAddr, err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Printf("Failed To Serve GRPC Server: %v", err)
	}
}
