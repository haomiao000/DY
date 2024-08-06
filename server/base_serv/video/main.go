package main

import (
	rpc_video "github.com/haomiao000/DY/internal/grpc_gen/rpc_video"
	api_client "github.com/haomiao000/DY/server/base_serv/video/api_client"
	api_server "github.com/haomiao000/DY/server/base_serv/video/api_server"
	configs "github.com/haomiao000/DY/server/base_serv/video/configs"
	dao "github.com/haomiao000/DY/server/base_serv/video/dao"
	initialize "github.com/haomiao000/DY/server/base_serv/video/initialize"

	grpc "google.golang.org/grpc"

	"fmt"
	"net"
)

func main() {
	db := initialize.InitDB()
	userServ := initialize.InitUser()
	grpcServer := grpc.NewServer()
	impl := &api_server.VideoServiceImpl{
		MysqlManager: dao.NewMysqlManager(db),
		UserManager:  api_client.NewUserClient(userServ),
	}
	rpc_video.RegisterVideoServiceImplServer(grpcServer, impl)

	listener, err := net.Listen("tcp", configs.VideoServerAddr)
	if err != nil {
		fmt.Printf("Failed To Listen On Addr %s: %v", configs.VideoServerAddr, err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Printf("Failed To Serve GRPC Server: %v", err)
	}
}
