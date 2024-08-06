package api_get

import (
	"context"
	"fmt"

	rpc_base "github.com/haomiao000/DY/internal/grpc_gen/rpc_base"
	rpc_video "github.com/haomiao000/DY/internal/grpc_gen/rpc_video"
)

type VideoManager struct {
	VideoService rpc_video.VideoServiceImplClient
}

func (s *VideoManager) UpdateVideoFavoriteCount(ctx context.Context, req *rpc_video.UpdateVideoFavoriteCountRequest) (*rpc_video.UpdateVideoFavoriteCountResponse, error) {
	res, err := s.VideoService.UpdateVideoFavoriteCount(ctx, req)
	if err != nil {
		fmt.Println("rpc serve error")
		return nil, err
	}
	return res, nil
}

func (s *VideoManager) GetFavoriteVideoListByVideoId(ctx context.Context, videoId []int64) ([]*rpc_base.Video, error) {
	res, err := s.VideoService.GetFavoriteVideoListByVideoId(ctx, &rpc_video.GetFavoriteVideoListByVideoIdRequest{
		VideoId: videoId,
	})
	if err != nil {
		fmt.Println("rpc serve error")
		return nil, err
	}
	return res.VideoList, nil
}

func NewVideoClient(client rpc_video.VideoServiceImplClient) *VideoManager {
	return &VideoManager{VideoService: client}
}

func (s *VideoManager) UpdateVideoCommentCount(ctx context.Context, req *rpc_video.UpdateVideoCommentCountRequest) (*rpc_video.UpdateVideoCommentCountResponse, error) {
	res, err := s.VideoService.UpdateVideoCommentCount(ctx, req)
	if err != nil {
		fmt.Println("rpc serve error")
		return nil, err
	}
	return res, nil
}
