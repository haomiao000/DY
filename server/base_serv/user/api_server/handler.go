package api_server

import (
	"context"
	"errors"
	"fmt"
	http "net/http"
	"strconv"

	rpc_base "github.com/haomiao000/DY/internal/grpc_gen/rpc_base"
	rpc_user "github.com/haomiao000/DY/internal/grpc_gen/rpc_user"
	model "github.com/haomiao000/DY/server/base_serv/user/model"
	configs "github.com/haomiao000/DY/server/common/configs"
	middleware "github.com/haomiao000/DY/server/common/middleware"

	snowflake "github.com/bwmarrin/snowflake"
	redis "github.com/haomiao000/DY/comm/redis"
	gorm "gorm.io/gorm"
)

type MysqlManager interface {
	CreateUserLoginInfo(userLoginInfo *model.UserLoginInfo) error
	CreateUser(user *model.User) error
	CheckUserLoginInfo(userLoginReq *rpc_user.UserLoginRequest) (*model.UserLoginInfo, error)
	GetUserByUid(userID int64) (*model.User, error)
	GetUserListByUserId(userID []int64) ([]*model.User, error)
}
type UserServiceImpl struct {
	rpc_user.UnimplementedUserServiceImplServer
	MysqlManager MysqlManager
}

func (s *UserServiceImpl) Register(ctx context.Context, req *rpc_user.UserRegisterRequest) (*rpc_user.UserRegisterResponse, error) {
	resp := new(rpc_user.UserRegisterResponse)
	sf, err := snowflake.NewNode(configs.UserSnowflakeNode)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  err.Error(),
		}
		return resp, err
	}
	user := &model.User{
		UserID:        sf.Generate().Int64(),
		Name:          req.Username,
		FollowCount:   0,
		FollowerCount: 0,
	}
	userLoginInfo := &model.UserLoginInfo{
		UserID:   user.UserID,
		Username: req.Username,
		Password: middleware.Gen_sha256(req.Password),
	}
	err = s.MysqlManager.CreateUserLoginInfo(userLoginInfo)
	if err == errors.New(configs.MysqlAlreadyExists) {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusConflict,
			StatusMsg:  "Username Already Exist",
		}
		return resp, err
	}
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  "Error Mysql Insert",
		}
		fmt.Println(resp)
		return resp, err
	}
	err = redis.SetJson(ctx , configs.LoginInfoRedisHead + req.Username , userLoginInfo)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  "Error Redis Insert",
		}
		return resp, err
	}
	err = s.MysqlManager.CreateUser(user)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  "Error Mysql Insert",
		}
		return resp, err
	}
	err = redis.SetJson(ctx , configs.UserRedisHead + strconv.FormatInt(user.UserID, 10) ,user)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  "Error Redis Insert User",
		}
		return resp, err
	}
	//返回的resp没用
	resp.BaseResp = &rpc_base.Response{

		StatusCode: http.StatusOK,
		StatusMsg:  "Successful Register User",
	}
	return resp, nil
}
func (s *UserServiceImpl) Login(ctx context.Context, req *rpc_user.UserLoginRequest) (*rpc_user.UserLoginResponse, error) {
	resp := new(rpc_user.UserLoginResponse)
	userLoginInfo := new(model.UserLoginInfo)
	exist , err := redis.GetJson(ctx , configs.LoginInfoRedisHead + req.Username , userLoginInfo)
	if err != nil || !exist {
		userLoginInfo , err = s.MysqlManager.CheckUserLoginInfo(req)
		if err == gorm.ErrRecordNotFound {
			resp.BaseResp = &rpc_base.Response{
				StatusCode: http.StatusNotFound,
				StatusMsg:  "Error Username or Password",
			}
			return resp, err
		}
		if err != nil {
			resp.BaseResp = &rpc_base.Response{
				StatusCode: http.StatusInternalServerError,
				StatusMsg:  "Error Mysql Serve",
			}
			return resp, err
		}
	}
	userID := userLoginInfo.UserID
	token, err := middleware.GenToken(userID)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  "Error Generate Token",
		}
		return resp, err
	}
	resp = &rpc_user.UserLoginResponse{
		BaseResp: &rpc_base.Response{
			StatusCode: http.StatusOK,
			StatusMsg:  "Successful Login",
		},
		UserId: userID,
		Token:  token,
	}
	return resp, nil
}
func (s *UserServiceImpl) GetUser(ctx context.Context, req *rpc_user.UserInfoRequest) (*rpc_user.UserResponse, error) {
	var resp = new(rpc_user.UserResponse)
	user := new(model.User)
	exist , err := redis.GetJson(ctx , configs.UserRedisHead + strconv.FormatInt(req.UserId, 10) , user)
	if err != nil || !exist {
		fmt.Println(err)
		user, err = s.MysqlManager.GetUserByUid(req.UserId)
		if err == gorm.ErrRecordNotFound {
			resp.BaseResp = &rpc_base.Response{
				StatusCode: http.StatusNotFound,
				StatusMsg:  "Error Found UserId",
			}
			return resp, err
		}
		if err != nil {
			resp.BaseResp = &rpc_base.Response{
				StatusCode: http.StatusInternalServerError,
				StatusMsg:  "Error Mysql Select",
			}
			return resp, err
		}
	}
	resp = &rpc_user.UserResponse{
		BaseResp: &rpc_base.Response{
			StatusCode: http.StatusOK,
			StatusMsg:  "Success Get User",
		},
		User: &rpc_base.User{
			Id:            user.UserID,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
		},
	}
	return resp, nil
}

func (s *UserServiceImpl) GetUserList(ctx context.Context, req *rpc_user.GetUserListRequest) (*rpc_user.GetUserListResponse, error) {
	var resp = new(rpc_user.GetUserListResponse)
	var userListStr []string
	for _ , o := range req.UserIdList {
		userListStr = append(userListStr , configs.UserRedisHead + strconv.FormatInt(o, 10))
	}
	var tmp []model.User
	userMap , err := redis.BatchGetJson(ctx , userListStr , tmp) 
	if err == nil {
		for _, v := range userMap {
			resp.UserList = append(resp.UserList, &rpc_base.User{
				Id:            v.(*model.User).UserID,
				Name:          v.(*model.User).Name,
				FollowCount:   v.(*model.User).FollowCount,
				FollowerCount: v.(*model.User).FollowerCount,
				IsFollow:      false,
			})
		}
		return resp, nil
	}
	userList, err := s.MysqlManager.GetUserListByUserId(req.UserIdList)
	if err != nil {
		return nil, err
	}
	for _, o := range userList {
		resp.UserList = append(resp.UserList, &rpc_base.User{
			Id:            o.UserID,
			Name:          o.Name,
			FollowCount:   o.FollowCount,
			FollowerCount: o.FollowerCount,
			IsFollow:      false,
		})
	}
	return resp, nil
}