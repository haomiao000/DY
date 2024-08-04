package handler

import (
	"context"
	"fmt"
	rpc_base "github.com/haomiao000/DY/server/grpc_gen/rpc_base"
	rpc_video "github.com/haomiao000/DY/server/grpc_gen/rpc_video"
	model "github.com/haomiao000/DY/server/service/video/model"
)

type MysqlManager interface {
	UpdateFavoriteCount(req *rpc_video.UpdateVideoFavoriteCountRequest) (error)
	GetFavoriteVideoList(req *rpc_video.GetFavoriteVideoListByVideoIdRequest) ([]*model.VideoRecord , error) 
	UpdateCommentCount(req *rpc_video.UpdateVideoCommentCountRequest) (error)
}

type VideoServiceImpl struct {
	rpc_video.UnimplementedVideoServiceImplServer
	MysqlManager
	UserManager
}
type UserManager interface {
	GetUser(ctx context.Context, userId int64) (*rpc_base.User, error)
}

func (s *VideoServiceImpl) UpdateVideoFavoriteCount(ctx context.Context, req *rpc_video.UpdateVideoFavoriteCountRequest) (*rpc_video.UpdateVideoFavoriteCountResponse, error) {
	var resp = new(rpc_video.UpdateVideoFavoriteCountResponse)
	err := s.MysqlManager.UpdateFavoriteCount(req)
	return resp , err
}

func (s *VideoServiceImpl) GetFavoriteVideoListByVideoId(ctx context.Context, req *rpc_video.GetFavoriteVideoListByVideoIdRequest) (*rpc_video.GetFavoriteVideoListByVideoIdResponse, error) {
	var resp = new(rpc_video.GetFavoriteVideoListByVideoIdResponse)
	res , err := s.MysqlManager.GetFavoriteVideoList(req)
	//这里不存在找不到的情况
	if err != nil {
		return nil , err
	}
	for _ , o := range res {
		auth , err := s.UserManager.GetUser(ctx , o.UserID)
		if err != nil {
			fmt.Printf("get user %d error\n",o.UserID)
			continue
		}
		resp.VideoList = append(resp.VideoList, &rpc_base.Video{
			Id: o.VideoID,
			Author: auth,
			PlayUrl: o.PlayUrl,
			CoverUrl: o.CoverUrl,
			FavoriteCount: o.FavoriteCount,
			CommentCount: o.CommentCount,
			IsFavorite: false,
		})
	}
	return resp , err
}

func (s *VideoServiceImpl) UpdateVideoCommentCount(ctx context.Context, req *rpc_video.UpdateVideoCommentCountRequest) (*rpc_video.UpdateVideoCommentCountResponse, error) {
	var resp = new(rpc_video.UpdateVideoCommentCountResponse)
	err := s.MysqlManager.UpdateCommentCount(req)
	return resp , err
}