package main

import (
	initialize "github.com/haomiao000/DY/server/service/user/initialize"
	handler "github.com/haomiao000/DY/server/service/user/handler"
	dao "github.com/haomiao000/DY/server/service/user/dao"
	grpc "google.golang.org/grpc"
	rpc_user "github.com/haomiao000/DY/server/grpc_gen/rpc_user"
	configs "github.com/haomiao000/DY/server/service/user/configs"
	"net"
	"fmt"
)

func main() {
	db := initialize.InitDB()
	grpcServer := grpc.NewServer()
	impl := &handler.UserServiceImpl{
		MysqlManager: dao.NewMysqlManager(db),
	}
	rpc_user.RegisterUserServiceImplServer(grpcServer , impl)

	listener, err := net.Listen("tcp", configs.UserServerAddr)
	if err != nil {
		fmt.Printf("Failed To Listen On Addr %s: %v", configs.UserServerAddr,err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Printf("Failed To Serve GRPC Server: %v", err)
	}
}