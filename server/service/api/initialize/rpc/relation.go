package rpc

import (
	relation "main/server/grpc_gen/rpc_relation"
	"main/server/service/api/configs"
	"google.golang.org/grpc"
)

func initRelation() {
	conn, err := grpc.Dial(configs.RelationServerAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	configs.GlobalRelationClient = relation.NewRelationServiceImplClient(conn)
}

