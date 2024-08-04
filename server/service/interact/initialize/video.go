package initialize 


import (
	video "main/server/grpc_gen/rpc_video"
	configs "main/server/service/interact/configs"
	grpc "google.golang.org/grpc"
)

func InitVideo() (video.VideoServiceImplClient) {
	conn, err := grpc.Dial(configs.VideoServerAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c := video.NewVideoServiceImplClient(conn)
	return c
}

