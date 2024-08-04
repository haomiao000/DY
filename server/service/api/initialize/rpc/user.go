package rpc

import (
	user "github.com/haomiao000/DY/server/grpc_gen/rpc_user"
	"github.com/haomiao000/DY/server/service/api/configs"
	"google.golang.org/grpc"
)

func initUser() {
	conn, err := grpc.Dial(configs.UserServerAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	configs.GlobalUserClient = user.NewUserServiceImplClient(conn)
}

