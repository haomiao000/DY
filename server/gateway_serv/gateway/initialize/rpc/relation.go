package rpc

import (
	relation "github.com/haomiao000/DY/internal/grpc_gen/rpc_relation"
	configs "github.com/haomiao000/DY/server/gateway_serv/gateway/configs"
	grpc "google.golang.org/grpc"
	discovery "github.com/haomiao000/DY/comm/discovery"
	insecure "google.golang.org/grpc/credentials/insecure"
)

func initRelation() {
	conn, err := grpc.NewClient("etcd:///relation", grpc.WithResolvers(discovery.GetResolver()),
	grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	configs.GlobalRelationClient = relation.NewRelationServiceImplClient(conn)
}
