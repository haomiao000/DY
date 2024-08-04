package main 

import (
	initialize "main/server/service/relation/initialize"
	handler "main/server/service/relation/handler"
	dao "main/server/service/relation/dao"
	grpc "google.golang.org/grpc"
	configs "main/server/service/relation/configs"
	rpc_relation "main/server/grpc_gen/rpc_relation"
	pkg "main/server/service/relation/pkg"
	"net"
	"fmt"
)

func main() {
	db := initialize.InitDB()
	userServ := initialize.InitUser()
	grpcServer := grpc.NewServer()
	impl := &handler.RelationServiceImpl{
		MysqlManager: dao.NewMysqlManager(db),
		UserManager: pkg.NewUserClient(userServ),
	}
	rpc_relation.RegisterRelationServiceImplServer(grpcServer , impl)

	listener, err := net.Listen("tcp", configs.RelationServerAddr)
	if err != nil {
		fmt.Printf("Failed To Listen On Addr %s: %v", configs.RelationServerAddr,err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Printf("Failed To Serve GRPC Server: %v", err)
	}
}