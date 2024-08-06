package rpc

import (
	interact "github.com/haomiao000/DY/internal/grpc_gen/rpc_interact"
	configs "github.com/haomiao000/DY/server/gateway_serv/gateway/configs"
	grpc "google.golang.org/grpc"
)

func initInteract() {
	conn, err := grpc.Dial(configs.InteractServerAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	configs.GlobalInteractClient = interact.NewInteractServiceImplClient(conn)
}
