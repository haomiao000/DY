package initialize 


import (
	user "github.com/haomiao000/DY/server/grpc_gen/rpc_user"
	configs "github.com/haomiao000/DY/server/service/video/configs"
	grpc "google.golang.org/grpc"
)

func InitUser() (user.UserServiceImplClient) {
	conn, err := grpc.Dial(configs.UserServerAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c := user.NewUserServiceImplClient(conn)
	return c
}

