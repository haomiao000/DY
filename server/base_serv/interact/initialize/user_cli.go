package initialize

import (
	user "github.com/haomiao000/DY/internal/grpc_gen/rpc_user"
	insecure "google.golang.org/grpc/credentials/insecure"
	discovery "github.com/haomiao000/DY/comm/discovery"
	grpc "google.golang.org/grpc"
	trace "github.com/haomiao000/DY/comm/trace"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	interceptor "github.com/haomiao000/DY/internal/interceptor"
)

func InitUser() user.UserServiceImplClient {
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
	c := user.NewUserServiceImplClient(conn)
	return c
}
