package pkg

import (
	rpc_user "github.com/haomiao000/DY/server/grpc_gen/rpc_user"
	rpc_base "github.com/haomiao000/DY/server/grpc_gen/rpc_base"
	"context"
	"fmt"
)

type UserManager struct {
	UserManager rpc_user.UserServiceImplClient
}

func (s *UserManager) GetUserList(ctx context.Context, userIds []int64) ([]*rpc_base.User, error) {
	res , err := s.UserManager.GetUserList(ctx , &rpc_user.GetUserListRequest{
		UserIdList: userIds,
	})
	if err != nil {
		fmt.Println("update rpc serve error")
		return nil , err
	}
	return res.UserList , nil
}

func NewUserClient(client rpc_user.UserServiceImplClient) *UserManager {
	return &UserManager{UserManager : client}
}