package api_client

import (
	"context"
	"fmt"

	"github.com/haomiao000/DY/comm/discovery"
	rpc_base "github.com/haomiao000/DY/internal/grpc_gen/rpc_base"
	rpc_user "github.com/haomiao000/DY/internal/grpc_gen/rpc_user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var userCli rpc_user.UserServiceImplClient

// 初始化 gRPC 客户端连接
func Init() error {
	con, err := grpc.NewClient("etcd:///video", grpc.WithResolvers(discovery.GetResolver()),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error in redis init")
		return err
	}
	userCli = rpc_user.NewUserServiceImplClient(con)
	return nil
}

func GetUser(ctx context.Context, userId []int64) (map[int64]*rpc_base.User, error) {
	_, err := userCli.GetUserList(ctx, &rpc_user.GetUserListRequest{
		UserIdList: userId,
	})

	if err != nil {
		fmt.Println("update rpc serve error")
		return nil, err
	}
	return map[int64]*rpc_base.User{}, nil
}
