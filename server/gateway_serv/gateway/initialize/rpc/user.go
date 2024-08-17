package rpc

import (
	discovery "github.com/haomiao000/DY/comm/discovery"
	user "github.com/haomiao000/DY/internal/grpc_gen/rpc_user"
	configs "github.com/haomiao000/DY/server/gateway_serv/gateway/configs"
	grpc "google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"
	trace "github.com/haomiao000/DY/comm/trace"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	interceptor "github.com/haomiao000/DY/internal/interceptor"
)

func initUser() {
	tracer , closer := trace.NewTracer("user")
	defer closer.Close()
	conn, err := grpc.NewClient(
		"etcd:///user", 
		grpc.WithResolvers(discovery.GetResolver()),
		grpc.WithUnaryInterceptor(
			grpcMiddleware.ChainUnaryClient(
				interceptor.ClientInterceptor(tracer),
			),
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	configs.GlobalUserClient = user.NewUserServiceImplClient(conn)
}
