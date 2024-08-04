package handler

import (
	configs "github.com/haomiao000/DY/server/common/configs"
	rpc_base "github.com/haomiao000/DY/server/grpc_gen/rpc_base"
	rpc_relation "github.com/haomiao000/DY/server/grpc_gen/rpc_relation"

	"context"
	http "net/http"

	"errors"
)

type MysqlManager interface {
	GetRelationStatus(req *rpc_relation.RelationActionRequest) (bool, error)
	CreateRelationInfo(userID int64 , toUserID int64) (error) 
	DeleteRelationInfo(userID int64 , toUserID int64) (error)
	GetFollowUserIdList(userID int64) ([]int64 , error)
	GetFollowerUserIdList(userID int64) ([]int64 , error)
	GetMutualFollowersIdList(userID int64) ([]int64 , error)
}
type UserManager interface {
	GetUserList(ctx context.Context, userIds []int64) ([]*rpc_base.User, error)
}
type RelationServiceImpl struct {
	rpc_relation.UnimplementedRelationServiceImplServer
	MysqlManager MysqlManager
	UserManager UserManager
}

func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *rpc_relation.RelationActionRequest) (*rpc_relation.RelationActionResponse, error) {
	var resp = new(rpc_relation.RelationActionResponse)
	if req.ActionType != configs.Follow && req.ActionType != configs.UnFollow {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusNotFound,
			StatusMsg: "Invalid Action Type",
		}
		return resp , errors.New("invalid action type")
	}
	if req.UserId == req.ToUserId {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusConflict,
			StatusMsg: "User Can Not Followe Itself",
		}
		return resp , errors.New("user can not followe itself")
	}
	followStatus , err := s.MysqlManager.GetRelationStatus(req)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg: "Error Get Relation Status",
		}
		return resp , err
	}
	if followStatus == configs.IsFollow {
		if req.ActionType == configs.Follow {
			resp.BaseResp = &rpc_base.Response{
				StatusCode: http.StatusOK,
				StatusMsg: "You Have Follow This User",
			}
			return resp , nil
		} else {
			err := s.MysqlManager.DeleteRelationInfo(req.UserId , req.ToUserId)
			if err != nil {
				resp.BaseResp = &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg: "Error Delete Relation",
				}
				return resp , err
			}
			resp.BaseResp = &rpc_base.Response{
				StatusCode: http.StatusOK,
				StatusMsg: "Successful Unfollow",
			}
			return resp , nil
		}
	}else {
		if req.ActionType == configs.UnFollow {
			resp.BaseResp = &rpc_base.Response{
				StatusCode: http.StatusOK,
				StatusMsg: "You Have Not Follow This User",
			}
			return resp , nil
		}else {
			err := s.MysqlManager.CreateRelationInfo(req.UserId , req.ToUserId)
			if err != nil {
				resp.BaseResp = &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg: "Error Create Relation",
				}
				return resp , err
			}
			resp.BaseResp = &rpc_base.Response{
				StatusCode: http.StatusOK,
				StatusMsg: "Successful Follow",
			}
			return resp , nil
		}
	}
}
func (s *RelationServiceImpl) GetFollowList(ctx context.Context, req *rpc_relation.RelationFollowListRequest) (*rpc_relation.RelationFollowListResponse, error) {
	var resp = new(rpc_relation.RelationFollowListResponse)
	mp , err := s.GetFollowMap(req.ViewerId)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg: "Error Get Follow Map",
		}
		return resp , err
	}
	followUserIdList , err := s.MysqlManager.GetFollowUserIdList(req.OwnerId)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg: "Error Get Follow User Id List",
		}
		return resp , err
	}
	followUserList , err := s.UserManager.GetUserList(ctx , followUserIdList)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg: "Error Get Follow User List",
		}
		return resp , err
	}
	for _ , o := range followUserList {
		resp.UserList = append(resp.UserList, &rpc_base.User{
			Id: o.Id,
			Name: o.Name,
			FollowCount: o.FollowCount,
			FollowerCount: o.FollowerCount,
			IsFollow: (*mp)[o.Id],
		})
	}
	resp.BaseResp = &rpc_base.Response{
		StatusCode: http.StatusOK,
		StatusMsg: "Successful Get Follow List",
	}
	return resp , nil
}
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *rpc_relation.RelationFollowerListRequest) (*rpc_relation.RelationFollowerListResponse, error) {
	var resp = new(rpc_relation.RelationFollowerListResponse)
	mp , err := s.GetFollowMap(req.ViewerId)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg: "Error Get Follow Map",
		}
		return resp , err
	}
	followerUserIdList , err := s.MysqlManager.GetFollowerUserIdList(req.OwnerId)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg: "Error Get Follower User Id List",
		}
		return resp , err
	}
	followerUserList , err := s.UserManager.GetUserList(ctx , followerUserIdList)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg: "Error Get Follower User List",
		}
		return resp , err
	}
	for _ , o := range followerUserList {
		resp.UserList = append(resp.UserList, &rpc_base.User{
			Id: o.Id,
			Name: o.Name,
			FollowCount: o.FollowCount,
			FollowerCount: o.FollowerCount,
			IsFollow: (*mp)[o.Id],
		})
	}
	resp.BaseResp = &rpc_base.Response{
		StatusCode: http.StatusOK,
		StatusMsg: "Successful Get Follower List",
	}
	return resp , nil
}
func (s *RelationServiceImpl) GetFriendList(ctx context.Context, req *rpc_relation.RelationFriendListRequest) (*rpc_relation.RelationFriendListResponse, error) {
	var resp = new(rpc_relation.RelationFriendListResponse)
	mp , err := s.GetFollowMap(req.ViewerId)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg: "Error Get Follow Map",
		}
		return resp , err
	}
	friendUserIdList , err := s.MysqlManager.GetMutualFollowersIdList(req.OwnerId)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg: "Error Get Friend User Id List",
		}
		return resp , err
	}
	friendUserList , err := s.UserManager.GetUserList(ctx , friendUserIdList)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg: "Error Get Friend User List",
		}
		return resp , err
	}
	for _ , o := range friendUserList {
		resp.UserList = append(resp.UserList, &rpc_base.User{
			Id: o.Id,
			Name: o.Name,
			FollowCount: o.FollowCount,
			FollowerCount: o.FollowerCount,
			IsFollow: (*mp)[o.Id],
		})
	}
	resp.BaseResp = &rpc_base.Response{
		StatusCode: http.StatusOK,
		StatusMsg: "Successful Get Follower List",
	}
	return resp , nil
}