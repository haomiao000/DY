package initialize

import (
	user "github.com/haomiao000/DY/internal/grpc_gen/rpc_user"
	grpc "google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"
	discovery "github.com/haomiao000/DY/comm/discovery"
)

func InitUser() user.UserServiceImplClient {
	conn, err := grpc.NewClient("etcd:///user", grpc.WithResolvers(discovery.GetResolver()),
	grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	c := user.NewUserServiceImplClient(conn)
	return c
}
