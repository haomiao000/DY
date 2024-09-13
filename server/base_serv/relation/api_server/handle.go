package api_server

import (
	"context"
	"errors"
	"fmt"
	http "net/http"
	"strconv"

	redis "github.com/haomiao000/DY/comm/redis"
	rpc_base "github.com/haomiao000/DY/internal/grpc_gen/rpc_base"
	rpc_relation "github.com/haomiao000/DY/internal/grpc_gen/rpc_relation"
	configs "github.com/haomiao000/DY/server/common/configs"
)

type MysqlManager interface {
	GetRelationStatus(req *rpc_relation.RelationActionRequest) (bool, error)
	CreateRelationInfo(userID int64, toUserID int64) error
	DeleteRelationInfo(userID int64, toUserID int64) error
	GetFollowUserIdList(userID int64) ([]int64, error)
	GetFollowerUserIdList(userID int64) ([]int64, error)
	GetMutualFollowersIdList(userID int64) ([]int64, error)
}
type UserManager interface {
	BatchGetUser(ctx context.Context, userIds []int64) (map[int64]*rpc_base.User, error)
}
type RelationServiceImpl struct {
	rpc_relation.UnimplementedRelationServiceImplServer
	MysqlManager MysqlManager
	UserManager  UserManager
}

func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *rpc_relation.RelationActionRequest) (*rpc_relation.RelationActionResponse, error) {
	var resp = new(rpc_relation.RelationActionResponse)
	if req.ActionType != configs.Follow && req.ActionType != configs.UnFollow {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusNotFound,
			StatusMsg:  "Invalid Action Type",
		}
		return resp, errors.New("invalid action type")
	}
	if req.UserId == req.ToUserId {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusConflict,
			StatusMsg:  "User Can Not Followe Itself",
		}
		return resp, errors.New("user can not followe itself")
	}
	followStatus, err := redis.SISMember(ctx , configs.UserFollowHead + strconv.FormatInt(req.UserId, 10) , strconv.FormatInt(req.ToUserId, 10))
	if err != nil || !followStatus {
		followStatus, err = s.MysqlManager.GetRelationStatus(req)
		if err != nil {
			resp.BaseResp = &rpc_base.Response{
				StatusCode: http.StatusInternalServerError,
				StatusMsg:  "Error Get Relation Status",
			}
			return resp, err
		}
	}
	if followStatus == configs.IsFollow {
		if req.ActionType == configs.Follow {
			resp.BaseResp = &rpc_base.Response{
				StatusCode: http.StatusOK,
				StatusMsg:  "You Have Follow This User",
			}
			return resp, nil
		} else {
			err := redis.SRem(ctx , configs.UserFollowHead + strconv.FormatInt(req.UserId, 10) , strconv.FormatInt(req.ToUserId, 10))
			if err != nil {
				resp.BaseResp = &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Delete Relation",
				}
				return resp, err
			}
			err = redis.SRem(ctx , configs.UserFollowerHead + strconv.FormatInt(req.ToUserId, 10) , strconv.FormatInt(req.UserId, 10))
			if err != nil {
				resp.BaseResp = &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Delete Relation",
				}
				return resp, err
			}
			err = s.MysqlManager.DeleteRelationInfo(req.UserId, req.ToUserId)
			if err != nil {
				resp.BaseResp = &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Delete Relation",
				}
				return resp, err
			}
			resp.BaseResp = &rpc_base.Response{
				StatusCode: http.StatusOK,
				StatusMsg:  "Successful Unfollow",
			}
			return resp, nil
		}
	} else {
		if req.ActionType == configs.UnFollow {
			resp.BaseResp = &rpc_base.Response{
				StatusCode: http.StatusOK,
				StatusMsg:  "You Have Not Follow This User",
			}
			return resp, nil
		} else {
			err := s.MysqlManager.CreateRelationInfo(req.UserId, req.ToUserId)
			if err != nil {
				resp.BaseResp = &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Create Relation",
				}
				return resp, err
			}
			err = redis.SAdd(ctx , configs.UserFollowHead + strconv.FormatInt(req.UserId, 10) , strconv.FormatInt(req.ToUserId, 10))
			if err != nil {
				resp.BaseResp = &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Create Relation",
				}
				return resp, err
			}
			err = redis.SAdd(ctx , configs.UserFollowerHead + strconv.FormatInt(req.ToUserId, 10) , strconv.FormatInt(req.UserId, 10))
			if err != nil {
				resp.BaseResp = &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Create Relation",
				}
				return resp, err
			}
			resp.BaseResp = &rpc_base.Response{
				StatusCode: http.StatusOK,
				StatusMsg:  "Successful Follow",
			}
			return resp, nil
		}
	}
}

func (s *RelationServiceImpl) GetFollowList(ctx context.Context, req *rpc_relation.RelationFollowListRequest) (*rpc_relation.RelationFollowListResponse, error) {
	var resp = new(rpc_relation.RelationFollowListResponse)
	mp, err := s.GetFollowMap(ctx , req.ViewerId)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  "Error Get Follow Map",
		}
		return resp, err
	}
	var followUserIdList []int64
	followUserIdListStr , err := redis.SMembers(ctx , configs.UserFollowHead + strconv.FormatInt(req.OwnerId, 10)) 
	if err != nil {
		followUserIdList, err = s.MysqlManager.GetFollowUserIdList(req.OwnerId)
		if err != nil {
			resp.BaseResp = &rpc_base.Response{
				StatusCode: http.StatusInternalServerError,
				StatusMsg:  "Error Get Follow User Id List",
			}
			return resp, err
		}
	}else {
		followUserIdList = make([]int64, len(followUserIdListStr))
		for i, str := range followUserIdListStr {
			user_id, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				resp.BaseResp = &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Get Follow User Id List",
				}
				return resp, err
			}
			followUserIdList[i] = user_id
		}
	}
	followUserMap, err := s.UserManager.BatchGetUser(ctx, followUserIdList)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  "Error Get Follow User List",
		}
		return resp, err
	}
	for _ , o := range followUserIdList {
		if _ , exist := followUserMap[o]; !exist {
			continue
		}
		resp.UserList = append(resp.UserList, &rpc_base.User{
			Id:            followUserMap[o].Id,
			Name:          followUserMap[o].Name,
			FollowCount:   followUserMap[o].FollowCount,
			FollowerCount: followUserMap[o].FollowerCount,
			IsFollow:      (*mp)[followUserMap[o].Id],
		})
	}
	resp.BaseResp = &rpc_base.Response{
		StatusCode: http.StatusOK,
		StatusMsg:  "Successful Get Follow List",
	}
	return resp, nil
}

func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *rpc_relation.RelationFollowerListRequest) (*rpc_relation.RelationFollowerListResponse, error) {
	var resp = new(rpc_relation.RelationFollowerListResponse)
	mp, err := s.GetFollowMap(ctx , req.ViewerId)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  "Error Get Follow Map",
		}
		return resp, err
	}
	var followerUserIdList []int64
	followerUserIdListStr , err := redis.SMembers(ctx , configs.UserFollowerHead + strconv.FormatInt(req.OwnerId, 10)) 
	fmt.Println(followerUserIdListStr)
	if err != nil {
		followerUserIdList, err = s.MysqlManager.GetFollowerUserIdList(req.OwnerId)
		if err != nil {
			resp.BaseResp = &rpc_base.Response{
				StatusCode: http.StatusInternalServerError,
				StatusMsg:  "Error Get Follower User Id List",
			}
			return resp, err
		}
	}else {
		followerUserIdList = make([]int64, len(followerUserIdListStr))
		for i, str := range followerUserIdListStr {
			user_id, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				resp.BaseResp = &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Get Follow User Id List",
				}
				return resp, err
			}
			followerUserIdList[i] = user_id
		}
	}
	followerUserMap, err := s.UserManager.BatchGetUser(ctx, followerUserIdList)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  "Error Get Follower User List",
		}
		return resp, err
	}
	for _ , o := range followerUserIdList {
		if _ , exist := followerUserMap[o]; !exist {
			continue
		}
		resp.UserList = append(resp.UserList, &rpc_base.User{
			Id:            followerUserMap[o].Id,
			Name:          followerUserMap[o].Name,
			FollowCount:   followerUserMap[o].FollowCount,
			FollowerCount: followerUserMap[o].FollowerCount,
			IsFollow:      (*mp)[followerUserMap[o].Id],
		})
	}
	resp.BaseResp = &rpc_base.Response{
		StatusCode: http.StatusOK,
		StatusMsg:  "Successful Get Follower List",
	}
	return resp, nil
}
func (s *RelationServiceImpl) GetFriendList(ctx context.Context, req *rpc_relation.RelationFriendListRequest) (*rpc_relation.RelationFriendListResponse, error) {
	var resp = new(rpc_relation.RelationFriendListResponse)
	mp, err := s.GetFollowMap(ctx , req.ViewerId)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  "Error Get Follow Map",
		}
		return resp, err
	}
	//这里没法做啊。。
	friendUserIdList, err := s.MysqlManager.GetMutualFollowersIdList(req.OwnerId)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  "Error Get Friend User Id List",
		}
		return resp, err
	}
	friendUserMap, err := s.UserManager.BatchGetUser(ctx, friendUserIdList)
	if err != nil {
		resp.BaseResp = &rpc_base.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  "Error Get Friend User List",
		}
		return resp, err
	}
	for _ , o := range friendUserIdList {
		if _ , exist := friendUserMap[o]; !exist {
			continue
		}
		resp.UserList = append(resp.UserList, &rpc_base.User{
			Id:            friendUserMap[o].Id,
			Name:          friendUserMap[o].Name,
			FollowCount:   friendUserMap[o].FollowCount,
			FollowerCount: friendUserMap[o].FollowerCount,
			IsFollow:      (*mp)[friendUserMap[o].Id],
		})
	}
	resp.BaseResp = &rpc_base.Response{
		StatusCode: http.StatusOK,
		StatusMsg:  "Successful Get Follower List",
	}
	return resp, nil
}
