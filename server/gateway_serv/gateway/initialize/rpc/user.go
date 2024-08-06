package rpc

import (
	user "github.com/haomiao000/DY/internal/grpc_gen/rpc_user"
	configs "github.com/haomiao000/DY/server/gateway_serv/gateway/configs"
	grpc "google.golang.org/grpc"
)

func initUser() {
	conn, err := grpc.Dial(configs.UserServerAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	configs.GlobalUserClient = user.NewUserServiceImplClient(conn)
}
