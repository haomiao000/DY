package initialize

import (
	video "github.com/haomiao000/DY/internal/grpc_gen/rpc_video"
	grpc "google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"
	discovery "github.com/haomiao000/DY/comm/discovery"
	trace "github.com/haomiao000/DY/comm/trace"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	interceptor "github.com/haomiao000/DY/internal/interceptor"
)

func InitVideo() video.VideoServiceImplClient {
	tracer , closer := trace.NewTracer("video")
	defer closer.Close()
	conn, err := grpc.NewClient(
		"etcd:///video", 
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
	c := video.NewVideoServiceImplClient(conn)
	return c
}
