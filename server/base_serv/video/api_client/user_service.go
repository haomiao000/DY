package api_client

import (
	"context"

	rpc_base "github.com/haomiao000/DY/internal/grpc_gen/rpc_base"
)

// var interactCli rpc_interact.InteractServiceImplClient

// var userCli rpc_user.UserServiceImplClient

// // 初始化 gRPC 客户端连接
// func Init() error {
// 	con, err := grpc.NewClient("etcd:///user", grpc.WithResolvers(discovery.GetResolver()),
// 		grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		fmt.Println("error in user init")
// 		return err
// 	}
// 	userCli = rpc_user.NewUserServiceImplClient(con)
// 	cc, err := grpc.NewClient("etcd:///interact", grpc.WithResolvers(discovery.GetResolver()),
// 		grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		fmt.Println("error in interact init")
// 		return err
// 	}
// 	interactCli = rpc_interact.NewInteractServiceImplClient(cc)
// 	return nil
// }

func GetUser(ctx context.Context, userId []int64) (map[int64]*rpc_base.User, error) {
	return map[int64]*rpc_base.User{}, nil
}

func GetUserFavoriteVideo(ctx context.Context, userID int64) (map[int64]bool, error) {
	return map[int64]bool{}, nil
	// rsp, err := interactCli.GetFavoriteVideoList(ctx, &rpc_interact.FavoriteListRequest{
	// 	OwnerId: userID,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// m := map[int64]bool{}
	// for _, videoID := range rsp.GetVideoList() {
	// 	m[videoID.Id] = true
	// }
	// return m, nil
}
