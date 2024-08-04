package main

import (
	initialize "main/server/service/user/initialize"
	handler "main/server/service/user/handler"
	dao "main/server/service/user/dao"
	grpc "google.golang.org/grpc"
	rpc_user "main/server/grpc_gen/rpc_user"
	configs "main/server/service/user/configs"
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