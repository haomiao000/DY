package rpc

import (
	interact "github.com/haomiao000/DY/internal/grpc_gen/rpc_interact"
	configs "github.com/haomiao000/DY/server/gateway_serv/gateway/configs"
	grpc "google.golang.org/grpc"
	discovery "github.com/haomiao000/DY/comm/discovery"
	insecure "google.golang.org/grpc/credentials/insecure"
	trace "github.com/haomiao000/DY/comm/trace"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	interceptor "github.com/haomiao000/DY/internal/interceptor"
)

func initInteract() {
	tracer , closer := trace.NewTracer("interact")
	defer closer.Close()
	conn, err := grpc.NewClient(
		"etcd:///interact", 
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
	configs.GlobalInteractClient = interact.NewInteractServiceImplClient(conn)
}
