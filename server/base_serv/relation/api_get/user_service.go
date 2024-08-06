package pkg

import (
	"context"
	"fmt"

	rpc_base "github.com/haomiao000/DY/internal/grpc_gen/rpc_base"
	rpc_user "github.com/haomiao000/DY/internal/grpc_gen/rpc_user"
)

type UserManager struct {
	UserManager rpc_user.UserServiceImplClient
}

func (s *UserManager) GetUserList(ctx context.Context, userIds []int64) ([]*rpc_base.User, error) {
	res, err := s.UserManager.GetUserList(ctx, &rpc_user.GetUserListRequest{
		UserIdList: userIds,
	})
	if err != nil {
		fmt.Println("update rpc serve error")
		return nil, err
	}
	return res.UserList, nil
}

func NewUserClient(client rpc_user.UserServiceImplClient) *UserManager {
	return &UserManager{UserManager: client}
}
