package rpc

import (
	discovery "github.com/haomiao000/DY/comm/discovery"
	user "github.com/haomiao000/DY/internal/grpc_gen/rpc_user"
	configs "github.com/haomiao000/DY/server/gateway_serv/gateway/configs"
	grpc "google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"
)

func initUser() {
	conn, err := grpc.NewClient("etcd:///user", grpc.WithResolvers(discovery.GetResolver()),
	grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	configs.GlobalUserClient = user.NewUserServiceImplClient(conn)

}
