package rpc

import (
	// discovery "github.com/haomiao000/DY/comm/discovery"
	user "github.com/haomiao000/DY/internal/grpc_gen/rpc_user"
	configs "github.com/haomiao000/DY/server/gateway_serv/gateway/configs"
	grpc "google.golang.org/grpc"
	// insecure "google.golang.org/grpc/credentials/insecure"
	otelgrpc "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

func initUser() {
	conn, err := grpc.Dial("127.0.0.1:8001", grpc.WithInsecure(), grpc.WithBlock(),
	grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		panic(err)
	}
	configs.GlobalUserClient = user.NewUserServiceImplClient(conn)
}
