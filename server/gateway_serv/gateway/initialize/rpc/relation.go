package rpc

import (
	relation "github.com/haomiao000/DY/internal/grpc_gen/rpc_relation"
	configs "github.com/haomiao000/DY/server/gateway_serv/gateway/configs"
	grpc "google.golang.org/grpc"
)

func initRelation() {
	conn, err := grpc.Dial(configs.RelationServerAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	configs.GlobalRelationClient = relation.NewRelationServiceImplClient(conn)
}
