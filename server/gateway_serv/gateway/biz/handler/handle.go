package api_server

import (
	"context"
	"errors"
	http "net/http"
	"time"

	gin "github.com/gin-gonic/gin"
	rpc_interact "github.com/haomiao000/DY/internal/grpc_gen/rpc_interact"
	rpc_relation "github.com/haomiao000/DY/internal/grpc_gen/rpc_relation"
	rpc_user "github.com/haomiao000/DY/internal/grpc_gen/rpc_user"
	configs "github.com/haomiao000/DY/server/gateway_serv/gateway/configs"
	model "github.com/haomiao000/DY/server/gateway_serv/gateway/model"
)

// Register .
// @router /douyin/user/register/ [POST]
func Register(c *gin.Context) {
	ctx := context.Background()
	ctx, fun := context.WithTimeout(ctx, time.Second*1)
	defer fun()
	var userRegisterReq model.UserRegisterRequest
	if err := c.ShouldBind(&userRegisterReq); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	res, err := configs.GlobalUserClient.Register(ctx, &rpc_user.UserRegisterRequest{
		Username: userRegisterReq.Username,
		Password: userRegisterReq.Password,
	})
	if err != nil {
		c.JSON(http.StatusGatewayTimeout, &model.UserRegisterResponse{
			BaseResp: &model.Response{
				StatusCode: 504,
				StatusMsg:  "time out",
			},
		})
		return
	}
	var resp = new(model.UserRegisterResponse)
	resp.BaseResp = &model.Response{
		StatusCode: res.BaseResp.StatusCode,
		StatusMsg:  res.BaseResp.StatusMsg,
	}
	c.JSON(http.StatusOK, resp)
}

// Login .
// @router /douyin/user/login/ [POST]
func Login(c *gin.Context) {
	o, _ := c.Get("ctx")
	ctx := o.(context.Context)
	var userLoginReq model.UserLoginRequest
	if err := c.ShouldBind(&userLoginReq); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	res, err := configs.GlobalUserClient.Login(ctx, &rpc_user.UserLoginRequest{
		Username: userLoginReq.Username,
		Password: userLoginReq.Password,
	})
	var resp = new(model.UserLoginResponse)
	resp.BaseResp = &model.Response{
		StatusCode: res.BaseResp.StatusCode,
		StatusMsg:  res.BaseResp.StatusMsg,
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.UserId = res.UserId
	resp.Token = res.Token
	c.JSON(http.StatusOK, resp)
}

// GetUserInfo .
// @router /douyin/user/ [GET]
func UserInfo(c *gin.Context) {
	o, _ := c.Get("ctx")
	ctx := o.(context.Context)
	var userInfoRequest model.UserInfoRequest
	if err := c.ShouldBind(&userInfoRequest); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		c.String(http.StatusNotFound, errors.New("api context get user_id failed").Error())
		return
	}
	res, err := configs.GlobalUserClient.GetUser(ctx, &rpc_user.UserInfoRequest{
		UserId:   userID.(int64),
		ViewerId: userInfoRequest.UserID,
	})
	var resp = new(model.UserResponse)
	resp.BaseResp = &model.Response{
		StatusCode: res.BaseResp.StatusCode,
		StatusMsg:  res.BaseResp.StatusMsg,
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.User = &model.User{
		Id:            res.User.Id,
		Name:          res.User.Name,
		FollowCount:   res.User.FollowCount,
		FollowerCount: res.User.FollowerCount,
		IsFollow:      res.User.IsFollow,
	}
	c.JSON(http.StatusOK, resp)
}

// FavoriteAction .
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(c *gin.Context) {
	o, _ := c.Get("ctx")
	ctx := o.(context.Context)
	var favoriteActionRequest model.FavoriteActionRequest
	if err := c.ShouldBind(&favoriteActionRequest); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		c.String(http.StatusNotFound, errors.New("api context get user_id failed").Error())
		return
	}
	res, err := configs.GlobalInteractClient.FavoriteAction(ctx, &rpc_interact.FavoriteActionRequest{
		UserId:     userID.(int64),
		VideoId:    favoriteActionRequest.VideoID,
		ActionType: int32(favoriteActionRequest.ActionType),
	})
	var resp = new(model.FavoriteActionResponse)
	resp.StatusCode = res.StatusCode
	resp.StatusMsg = res.StatusMsg
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// FavoriteList .
// @router /douyin/favorite/list/ [GET]
func FavoriteList(c *gin.Context) {
	o, _ := c.Get("ctx")
	ctx := o.(context.Context)
	var favoriteListRequest model.FavoriteListRequest
	if err := c.ShouldBind(&favoriteListRequest); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		c.String(http.StatusNotFound, errors.New("api context get user_id failed").Error())
		return
	}
	res, err := configs.GlobalInteractClient.GetFavoriteVideoList(ctx, &rpc_interact.FavoriteListRequest{
		OwnerId:  favoriteListRequest.UserID,
		ViewerId: userID.(int64),
	})
	var resp = new(model.FavoriteListResponse)
	resp.StatusCode = res.StatusCode
	resp.StatusMsg = res.StatusMsg
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	for _, o := range res.VideoList {
		resp.VideoList = append(resp.VideoList, &model.Video{
			Id: o.Id,
			Author: &model.User{
				Id:            o.Author.Id,
				Name:          o.Author.Name,
				FollowCount:   o.Author.FollowCount,
				FollowerCount: o.Author.FollowerCount,
				IsFollow:      o.Author.IsFollow,
			},
			PlayUrl:       o.PlayUrl,
			CoverUrl:      o.CoverUrl,
			FavoriteCount: o.FavoriteCount,
			CommentCount:  o.CommentCount,
			IsFavorite:    o.IsFavorite,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// CommentAction .
// @router /douyin/comment/action/ [POST]
func CommentAction(c *gin.Context) {
	o, _ := c.Get("ctx")
	ctx := o.(context.Context)
	var commentActionRequest model.CommentActionRequest
	if err := c.ShouldBind(&commentActionRequest); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		c.String(http.StatusNotFound, errors.New("api context get user_id failed").Error())
		return
	}
	res, err := configs.GlobalInteractClient.CommentAction(ctx, &rpc_interact.CommentActionRequest{
		UserId:      userID.(int64),
		VideoId:     commentActionRequest.VideoID,
		ActionType:  int32(commentActionRequest.ActionType),
		CommentText: commentActionRequest.CommentText,
		CommentId:   commentActionRequest.CommentID,
	})
	var resp = new(model.CommentActionResponse)
	resp.BaseResp = &model.Response{
		StatusCode: res.BaseResp.StatusCode,
		StatusMsg:  res.BaseResp.StatusMsg,
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	if res.Comment != nil {
		resp.Comment = &model.Comment{
			Id: res.Comment.Id,
			User: &model.User{
				Id:            res.Comment.User.Id,
				Name:          res.Comment.User.Name,
				FollowCount:   res.Comment.User.FollowCount,
				FollowerCount: res.Comment.User.FollowerCount,
				IsFollow:      res.Comment.User.IsFollow,
			},
			Content:    res.Comment.Content,
			CreateDate: res.Comment.CreateDate,
		}
	}
	c.JSON(http.StatusOK, resp)
}

// CommentList .
// @router /douyin/comment/list/ [GET]
func CommentList(c *gin.Context) {
	o, _ := c.Get("ctx")
	ctx := o.(context.Context)
	var commentListRequest model.CommentListRequest
	if err := c.ShouldBind(&commentListRequest); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	res, err := configs.GlobalInteractClient.GetCommentList(ctx, &rpc_interact.CommentListRequest{
		VideoId: commentListRequest.VideoID,
	})
	var resp = new(model.CommentListResponse)
	resp.BaseResp = &model.Response{
		StatusCode: res.BaseResp.StatusCode,
		StatusMsg:  res.BaseResp.StatusMsg,
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	for _, o := range res.CommentList {
		resp.CommentList = append(resp.CommentList, &model.Comment{
			Id: o.Id,
			User: &model.User{
				Id:            o.User.Id,
				Name:          o.User.Name,
				FollowCount:   o.User.FollowCount,
				FollowerCount: o.User.FollowerCount,
				IsFollow:      o.User.IsFollow,
			},
			Content:    o.Content,
			CreateDate: o.CreateDate,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// Action .
// @router /douyin/relation/action/ [POST]
func RelationAction(c *gin.Context) {
	o, _ := c.Get("ctx")
	ctx := o.(context.Context)
	var relationActionRequest model.RelationActionRequest
	if err := c.ShouldBind(&relationActionRequest); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		c.String(http.StatusNotFound, errors.New("api context get user_id failed").Error())
		return
	}
	res, err := configs.GlobalRelationClient.RelationAction(ctx, &rpc_relation.RelationActionRequest{
		UserId:     userID.(int64),
		ToUserId:   relationActionRequest.ToUserId,
		ActionType: int32(relationActionRequest.ActionType),
	})
	var resp = new(model.RelationActionResponse)
	resp.BaseResp = &model.Response{
		StatusCode: res.BaseResp.StatusCode,
		StatusMsg:  res.BaseResp.StatusMsg,
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// FollowList .
// @router /douyin/relation/follow/list/ [GET]
func FollowList(c *gin.Context) {
	o, _ := c.Get("ctx")
	ctx := o.(context.Context)
	var relationFollowListRequest model.RelationFollowListRequest
	if err := c.ShouldBind(&relationFollowListRequest); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		c.String(http.StatusNotFound, errors.New("api context get user_id failed").Error())
		return
	}
	res, err := configs.GlobalRelationClient.GetFollowList(ctx, &rpc_relation.RelationFollowListRequest{
		ViewerId: userID.(int64),
		OwnerId:  relationFollowListRequest.UserID,
	})
	var resp = new(model.UserListResponse)
	resp.BaseResp = &model.Response{
		StatusCode: res.BaseResp.StatusCode,
		StatusMsg:  res.BaseResp.StatusMsg,
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	for _, o := range res.UserList {
		resp.UserList = append(resp.UserList, &model.User{
			Id:            o.Id,
			Name:          o.Name,
			FollowCount:   o.FollowCount,
			FollowerCount: o.FollowerCount,
			IsFollow:      o.IsFollow,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// FollowerList .
// @router /douyin/relation/follower/list/ [GET]
func FollowerList(c *gin.Context) {
	o, _ := c.Get("ctx")
	ctx := o.(context.Context)
	var relationFollowerListRequest model.RelationFollowerListRequest
	if err := c.ShouldBind(&relationFollowerListRequest); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		c.String(http.StatusNotFound, errors.New("api context get user_id failed").Error())
		return
	}
	res, err := configs.GlobalRelationClient.GetFollowerList(ctx, &rpc_relation.RelationFollowerListRequest{
		ViewerId: userID.(int64),
		OwnerId:  relationFollowerListRequest.UserID,
	})
	var resp = new(model.UserListResponse)
	resp.BaseResp = &model.Response{
		StatusCode: res.BaseResp.StatusCode,
		StatusMsg:  res.BaseResp.StatusMsg,
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	for _, o := range res.UserList {
		resp.UserList = append(resp.UserList, &model.User{
			Id:            o.Id,
			Name:          o.Name,
			FollowCount:   o.FollowCount,
			FollowerCount: o.FollowerCount,
			IsFollow:      o.IsFollow,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// FriendList .
// @router /douyin/relation/friend/list/ [GET]
func FriendList(c *gin.Context) {
	o, _ := c.Get("ctx")
	ctx := o.(context.Context)
	var relationFriendListRequest model.RelationFriendListRequest
	if err := c.ShouldBind(&relationFriendListRequest); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.String(http.StatusNotFound, errors.New("api context get user_id failed").Error())
		return
	}
	res, err := configs.GlobalRelationClient.GetFriendList(ctx, &rpc_relation.RelationFriendListRequest{
		ViewerId: userID.(int64),
		OwnerId:  relationFriendListRequest.UserID,
	})
	var resp = new(model.UserListResponse)
	resp.BaseResp = &model.Response{
		StatusCode: res.BaseResp.StatusCode,
		StatusMsg:  res.BaseResp.StatusMsg,
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	for _, o := range res.UserList {
		resp.UserList = append(resp.UserList, &model.User{
			Id:            o.Id,
			Name:          o.Name,
			FollowCount:   o.FollowCount,
			FollowerCount: o.FollowerCount,
			IsFollow:      o.IsFollow,
		})
	}
	c.JSON(http.StatusOK, resp)
}
