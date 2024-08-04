package rpc

import (
	interact "main/server/grpc_gen/rpc_interact"
	"main/server/service/api/configs"
	"google.golang.org/grpc"
)

func initInteract() {
	conn, err := grpc.Dial(configs.InteractServerAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	configs.GlobalInteractClient = interact.NewInteractServiceImplClient(conn)
}

