package api_get

import (
	"context"
	"fmt"

	rpc_base "github.com/haomiao000/DY/internal/grpc_gen/rpc_base"
	rpc_user "github.com/haomiao000/DY/internal/grpc_gen/rpc_user"
)

type UserManager struct {
	UserManager rpc_user.UserServiceImplClient
}

func (s *UserManager) GetUser(ctx context.Context, userId int64) (*rpc_base.User, error) {
	res, err := s.UserManager.GetUser(ctx, &rpc_user.UserInfoRequest{
		UserId: userId,
	})
	if err != nil {
		fmt.Println("update rpc serve error")
		return nil, err
	}
	return res.User, nil
}

func NewUserClient(client rpc_user.UserServiceImplClient) *UserManager {
	return &UserManager{UserManager: client}
}
