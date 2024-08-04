package main 

import (
	initialize "github.com/haomiao000/DY/server/service/relation/initialize"
	handler "github.com/haomiao000/DY/server/service/relation/handler"
	dao "github.com/haomiao000/DY/server/service/relation/dao"
	grpc "google.golang.org/grpc"
	configs "github.com/haomiao000/DY/server/service/relation/configs"
	rpc_relation "github.com/haomiao000/DY/server/grpc_gen/rpc_relation"
	pkg "github.com/haomiao000/DY/server/service/relation/pkg"
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