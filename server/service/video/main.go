package main

import (
	rpc_video "main/server/grpc_gen/rpc_video"
	configs "main/server/service/video/configs"
	dao "main/server/service/video/dao"
	handler "main/server/service/video/handler"
	initialize "main/server/service/video/initialize"
	pkg "main/server/service/video/pkg"

	grpc "google.golang.org/grpc"

	"fmt"
	"net"
)

func main() {
	db := initialize.InitDB()
	userServ := initialize.InitUser()
	grpcServer := grpc.NewServer()
	impl := &handler.VideoServiceImpl{
		MysqlManager: dao.NewMysqlManager(db),
		UserManager: pkg.NewUserClient(userServ),
	}
	rpc_video.RegisterVideoServiceImplServer(grpcServer , impl)

	listener, err := net.Listen("tcp", configs.VideoServerAddr)
	if err != nil {
		fmt.Printf("Failed To Listen On Addr %s: %v", configs.VideoServerAddr,err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Printf("Failed To Serve GRPC Server: %v", err)
	}
}