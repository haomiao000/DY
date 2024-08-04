package rpc

import (
	user "main/server/grpc_gen/rpc_user"
	"main/server/service/api/configs"
	"google.golang.org/grpc"
)

func initUser() {
	conn, err := grpc.Dial(configs.UserServerAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	configs.GlobalUserClient = user.NewUserServiceImplClient(conn)
}

