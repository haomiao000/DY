package initialize

import (
	video "github.com/haomiao000/DY/internal/grpc_gen/rpc_video"
	configs "github.com/haomiao000/DY/server/base_serv/interact/configs"
	grpc "google.golang.org/grpc"
)

func InitVideo() video.VideoServiceImplClient {
	conn, err := grpc.Dial(configs.VideoServerAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c := video.NewVideoServiceImplClient(conn)
	return c
}
