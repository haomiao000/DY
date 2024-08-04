package initialize 


import (
	user "main/server/grpc_gen/rpc_user"
	configs "main/server/service/interact/configs"
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

