package main

import (
	"fmt"
	"net"

	rpc_relation "github.com/haomiao000/DY/internal/grpc_gen/rpc_relation"
	api_get "github.com/haomiao000/DY/server/base_serv/relation/api_get"
	api_set "github.com/haomiao000/DY/server/base_serv/relation/api_set"
	configs "github.com/haomiao000/DY/server/base_serv/relation/configs"
	dao "github.com/haomiao000/DY/server/base_serv/relation/dao"
	initialize "github.com/haomiao000/DY/server/base_serv/relation/initialize"
	grpc "google.golang.org/grpc"
)

func main() {
	db := initialize.InitDB()
	userServ := initialize.InitUser()
	grpcServer := grpc.NewServer()
	impl := &api_set.RelationServiceImpl{
		MysqlManager: dao.NewMysqlManager(db),
		UserManager:  api_get.NewUserClient(userServ),
	}
	rpc_relation.RegisterRelationServiceImplServer(grpcServer, impl)

	listener, err := net.Listen("tcp", configs.RelationServerAddr)
	if err != nil {
		fmt.Printf("Failed To Listen On Addr %s: %v", configs.RelationServerAddr, err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Printf("Failed To Serve GRPC Server: %v", err)
	}
}
