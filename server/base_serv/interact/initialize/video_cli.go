package initialize

import (
	video "github.com/haomiao000/DY/internal/grpc_gen/rpc_video"
	grpc "google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"
	discovery "github.com/haomiao000/DY/comm/discovery"
)

func InitVideo() video.VideoServiceImplClient {
	conn, err := grpc.NewClient("etcd:///video", grpc.WithResolvers(discovery.GetResolver()),
	grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	c := video.NewVideoServiceImplClient(conn)
	return c
}
