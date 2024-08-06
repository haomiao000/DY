package main

import (
	"fmt"
	"net"

	rpc_interact "github.com/haomiao000/DY/internal/grpc_gen/rpc_interact"
	api_get "github.com/haomiao000/DY/server/base_serv/interact/api_get"
	api_set "github.com/haomiao000/DY/server/base_serv/interact/api_set"
	configs "github.com/haomiao000/DY/server/base_serv/interact/configs"
	dao "github.com/haomiao000/DY/server/base_serv/interact/dao"
	initialize "github.com/haomiao000/DY/server/base_serv/interact/initialize"
	grpc "google.golang.org/grpc"
)

func main() {
	db := initialize.InitDB()
	userServ := initialize.InitUser()
	videoServ := initialize.InitVideo()
	grpcServer := grpc.NewServer()
	impl := &api_set.InteractServiceImpl{
		FavoriteMysqlManager: dao.NewMysqlManager(db),
		CommentMysqlManager:  dao.NewMysqlManager(db),
		UserManager:          api_get.NewUserClient(userServ),
		VideoManager:         api_get.NewVideoClient(videoServ),
	}
	rpc_interact.RegisterInteractServiceImplServer(grpcServer, impl)

	listener, err := net.Listen("tcp", configs.InteractServerAddr)
	if err != nil {
		fmt.Printf("Failed To Listen On Addr %s: %v", configs.InteractServerAddr, err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Printf("Failed To Serve GRPC Server: %v", err)
	}
}
