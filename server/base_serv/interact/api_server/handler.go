package api_server

import (
	"context"
	"errors"
	"fmt"

	rpc_base "github.com/haomiao000/DY/internal/grpc_gen/rpc_base"
	rpc_interact "github.com/haomiao000/DY/internal/grpc_gen/rpc_interact"
	rpc_video "github.com/haomiao000/DY/internal/grpc_gen/rpc_video"
	model "github.com/haomiao000/DY/server/base_serv/interact/model"
	configs "github.com/haomiao000/DY/server/common/configs"

	http "net/http"
	"strconv"
	"time"

	snowflake "github.com/bwmarrin/snowflake"
	redis "github.com/haomiao000/DY/comm/redis"
	gorm "gorm.io/gorm"
)

type FavoriteMysqlManager interface {
	GetFavoriteStatus(req *rpc_interact.FavoriteActionRequest) (bool, error)
	CreateFavorite(favorite *model.Favorite) error
	DeleteFavorite(req *rpc_interact.FavoriteActionRequest) error
	GetFavoriteVideoIdList(req *rpc_interact.FavoriteListRequest) ([]int64, error)
}
type CommentMysqlManager interface {
	CreateComment(comment *model.Comment) error
	DeleteComment(commentID int64) error
	GetCommentList(videoID int64) ([]*model.Comment, error)
}
type VideoManager interface {
	UpdateVideoFavoriteCount(ctx context.Context, req *rpc_video.UpdateVideoFavoriteCountRequest) (*rpc_video.UpdateVideoFavoriteCountResponse, error)
	GetFavoriteVideoListByVideoId(ctx context.Context, videoId []int64) ([]*rpc_base.Video, error)
	UpdateVideoCommentCount(ctx context.Context, req *rpc_video.UpdateVideoCommentCountRequest) (*rpc_video.UpdateVideoCommentCountResponse, error)
}
type UserManager interface {
	GetUser(ctx context.Context, userId int64) (*rpc_base.User, error)
}
type InteractServiceImpl struct {
	rpc_interact.UnimplementedInteractServiceImplServer
	FavoriteMysqlManager FavoriteMysqlManager
	CommentMysqlManager  CommentMysqlManager
	UserManager          UserManager
	VideoManager         VideoManager
}

func (s *InteractServiceImpl) FavoriteAction(ctx context.Context, req *rpc_interact.FavoriteActionRequest) (*rpc_interact.FavoriteActionResponse, error) {
	resp := new(rpc_interact.FavoriteActionResponse)
	favoriteStatus, err := redis.SISMember(ctx , configs.FavoriteOwnerHead + strconv.FormatInt(req.UserId, 10) , strconv.FormatInt(req.VideoId, 10))
	if err != nil || !favoriteStatus {
		favoriteStatus, err = s.FavoriteMysqlManager.GetFavoriteStatus(req)
		if err != nil {
			resp = &rpc_interact.FavoriteActionResponse{
				StatusCode: http.StatusInternalServerError,
				StatusMsg:  "Error Get Favorite Status",
			}
			return resp, err
		}
	}
	if req.ActionType != configs.Like && req.ActionType != configs.UnLike {
		resp.StatusCode = http.StatusNotFound
		resp.StatusMsg = "Invalid Action Type"
		return resp, errors.New("invalid action type")
	}
	if favoriteStatus == configs.IsLike {
		if req.ActionType == configs.Like {
			resp = &rpc_interact.FavoriteActionResponse{
				StatusCode: http.StatusOK,
				StatusMsg:  "You Like The Video You Like",
			}
			return resp, nil
		} else {
			err := redis.SRem(ctx , configs.FavoriteOwnerHead + strconv.FormatInt(req.UserId, 10) , strconv.FormatInt(req.VideoId, 10))
			if err != nil {
				resp = &rpc_interact.FavoriteActionResponse{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Delete Favorite",
				}
				return resp, err
			}
			err = s.FavoriteMysqlManager.DeleteFavorite(req)
				if err != nil {
					resp = &rpc_interact.FavoriteActionResponse{
						StatusCode: http.StatusInternalServerError,
						StatusMsg:  "Error Delete Favorite",
					}
					return resp, err
				}
			//根据视频格式更新，视频存储为set
			_, err = s.VideoManager.UpdateVideoFavoriteCount(ctx, &rpc_video.UpdateVideoFavoriteCountRequest{
				VideoId:      req.VideoId,
				ChangeNumber: configs.Minus_like,
			})
			if err != nil {
				resp = &rpc_interact.FavoriteActionResponse{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Update Video Favorite Count",
				}
				return resp, err
			}
			resp = &rpc_interact.FavoriteActionResponse{
				StatusCode: http.StatusOK,
				StatusMsg:  "Successful Delete Favrite",
			}
			return resp, nil
		}
	} else {
		if req.ActionType == configs.UnLike {
			resp = &rpc_interact.FavoriteActionResponse{
				StatusCode: http.StatusOK,
				StatusMsg:  "You UnLike The Video You UnLike",
			}
			return resp, nil
		} else {
			favorite := &model.Favorite{
				UserID:     req.UserId,
				VideoID:    req.VideoId,
				CreateDate: time.Now().UnixNano(),
			}
			if err := s.FavoriteMysqlManager.CreateFavorite(favorite); err != nil {
				resp = &rpc_interact.FavoriteActionResponse{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Create Favorite",
				}
				return resp, err
			}
			if err := redis.SAdd(ctx , configs.FavoriteOwnerHead + strconv.FormatInt(favorite.UserID, 10) , strconv.FormatInt(favorite.VideoID, 10)); err != nil {
				resp = &rpc_interact.FavoriteActionResponse{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Create Favorite",
				}
				return resp, err
			}
			_, err = s.VideoManager.UpdateVideoFavoriteCount(ctx, &rpc_video.UpdateVideoFavoriteCountRequest{
				VideoId:      req.VideoId,
				ChangeNumber: configs.Plus_like,
			})
			if err != nil {
				resp = &rpc_interact.FavoriteActionResponse{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Update Video Favorite Count",
				}
				return resp, err
			}
			resp = &rpc_interact.FavoriteActionResponse{
				StatusCode: http.StatusOK,
				StatusMsg:  "Successful Create Favrite",
			}
			return resp, nil
		}
	}
}
func (s *InteractServiceImpl) GetFavoriteVideoList(ctx context.Context, req *rpc_interact.FavoriteListRequest) (*rpc_interact.FavoriteListResponse, error) {
	resp := new(rpc_interact.FavoriteListResponse)
	var favoriteVideoIdList []int64
	favoriteVideoIdListStr , err := redis.SMembers(ctx , configs.FavoriteOwnerHead + strconv.FormatInt(req.OwnerId, 10))
	if favoriteVideoIdListStr == nil && err == nil{
		resp = &rpc_interact.FavoriteListResponse{
			StatusCode: http.StatusOK,
			StatusMsg:  "Successful Get Video List",
			VideoList:  nil,
		}
		return resp, nil
	}
	if err != nil {
		favoriteVideoIdList, err = s.FavoriteMysqlManager.GetFavoriteVideoIdList(req)
		if err == gorm.ErrRecordNotFound {
			resp = &rpc_interact.FavoriteListResponse{
				StatusCode: http.StatusOK,
				StatusMsg:  "Successful Get Video List",
				VideoList:  nil,
			}
			return resp, nil
		}
		if err != nil {
			resp = &rpc_interact.FavoriteListResponse{
				StatusCode: http.StatusInternalServerError,
				StatusMsg:  "Error Get Video Id List",
				VideoList:  nil,
			}
			return resp, err
		}
	}else {
		favoriteVideoIdList = make([]int64, len(favoriteVideoIdListStr))
		for i, str := range favoriteVideoIdListStr {
			videoId, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				// Handle the conversion error if needed
				resp = &rpc_interact.FavoriteListResponse{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "failed to parse video ID",
					VideoList:  nil,
				}
				return resp, err
			}
			favoriteVideoIdList[i] = videoId
		}
	}
	favoriteVideoList, err := s.VideoManager.GetFavoriteVideoListByVideoId(ctx, favoriteVideoIdList)
	if err != nil {
		resp = &rpc_interact.FavoriteListResponse{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  "Error Get Video List",
			VideoList:  nil,
		}
		return resp, err
	}
	resp = &rpc_interact.FavoriteListResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  "Successful Get Video List",
		VideoList:  favoriteVideoList,
	}
	return resp, nil
}
func (s *InteractServiceImpl) CommentAction(ctx context.Context, req *rpc_interact.CommentActionRequest) (*rpc_interact.CommentActionResponse, error) {
	var resp = new(rpc_interact.CommentActionResponse)
	if req.ActionType == configs.AddComment {
		tmp, err := snowflake.NewNode(configs.CommentSnowFlakeNode)
		if err != nil {
			resp = &rpc_interact.CommentActionResponse{
				BaseResp: &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Generate Comment ID Flake Node",
				},
			}
			return resp, err
		}
		comment := &model.Comment{
			CommentID:   tmp.Generate().Int64(),
			UserId:      req.UserId,
			VideoId:     req.VideoId,
			ActionType:  int8(req.ActionType),
			CommentText: req.CommentText,
			CreateDate:  time.Now().Unix(),
		}
		err = s.CommentMysqlManager.CreateComment(comment)
		if err != nil {
			resp = &rpc_interact.CommentActionResponse{
				BaseResp: &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Create Comment",
				},
			}
			return resp, err
		}
		err = redis.SAdd(ctx , configs.VideoCommentHead + strconv.FormatInt(comment.VideoId, 10) , strconv.FormatInt(comment.CommentID, 10))
		if err != nil {
			resp = &rpc_interact.CommentActionResponse{
				BaseResp: &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Create Comment In Redis",
				},
			}
			return resp, err
		}
		err = redis.SetJson(ctx , configs.SingleCommentHead + strconv.FormatInt(comment.CommentID, 10) , comment)
		if err != nil {
			resp = &rpc_interact.CommentActionResponse{
				BaseResp: &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Create Comment In Redis",
				},
			}
			return resp, err
		}
		_, err = s.VideoManager.UpdateVideoCommentCount(ctx, &rpc_video.UpdateVideoCommentCountRequest{
			VideoId:      req.VideoId,
			ChangeNumber: configs.Plus_comment,
		})
		if err != nil {
			resp = &rpc_interact.CommentActionResponse{
				BaseResp: &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Update Comment Count",
				},
			}
			return resp, err
		}
		auth, err := s.UserManager.GetUser(ctx, req.UserId)
		if err != nil {
			resp = &rpc_interact.CommentActionResponse{
				BaseResp: &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Get Comment Author",
				},
			}
			return resp, err
		}
		resp = &rpc_interact.CommentActionResponse{
			BaseResp: &rpc_base.Response{
				StatusCode: http.StatusOK,
				StatusMsg:  "Successful Create Comment",
			},
			Comment: &rpc_base.Comment{
				Id:         comment.CommentID,
				User:       auth,
				Content:    comment.CommentText,
				CreateDate: strconv.FormatInt(comment.CreateDate, 10),
			},
		}
		return resp, nil
	} else if req.ActionType == configs.DeleteComment {
		err := redis.SRem(ctx , configs.VideoCommentHead + strconv.FormatInt(req.VideoId, 10) , strconv.FormatInt(req.CommentId, 10))
		if err != nil {
			resp = &rpc_interact.CommentActionResponse{
				BaseResp: &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Delete Comment In Redis",
				},
			}
			return resp, err
		}
		exist , err := redis.Delete(ctx , configs.SingleCommentHead + strconv.FormatInt(req.CommentId, 10))
		if err != nil || !exist {
			resp = &rpc_interact.CommentActionResponse{
				BaseResp: &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Delete Comment In Redis",
				},
			}
			return resp, err
		}
		err = s.CommentMysqlManager.DeleteComment(req.CommentId)
		if err != nil {
			resp = &rpc_interact.CommentActionResponse{
				BaseResp: &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Delete Comment",
				},
			}
			return resp, err
		}
		_, err = s.VideoManager.UpdateVideoCommentCount(ctx, &rpc_video.UpdateVideoCommentCountRequest{
			VideoId:      req.VideoId,
			ChangeNumber: configs.Minus_comment,
		})
		if err != nil {
			resp = &rpc_interact.CommentActionResponse{
				BaseResp: &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Update Comment Count",
				},
			}
			return resp, err
		}
		resp = &rpc_interact.CommentActionResponse{
			BaseResp: &rpc_base.Response{
				StatusCode: http.StatusOK,
				StatusMsg:  "Successful Delete Comment Count",
			},
		}
		return resp, nil
	} else {
		resp = &rpc_interact.CommentActionResponse{
			BaseResp: &rpc_base.Response{
				StatusCode: http.StatusNotFound,
				StatusMsg:  "Invalid Comment ActionType",
			},
		}
		return resp, errors.New("invalid comment actionType")
	}
}
func (s *InteractServiceImpl) GetCommentList(ctx context.Context, req *rpc_interact.CommentListRequest) (*rpc_interact.CommentListResponse, error) {
	var resp = new(rpc_interact.CommentListResponse)
	var commentList []*model.Comment
	commentIdListStr , err := redis.SMembers(ctx , configs.VideoCommentHead + strconv.FormatInt(req.VideoId, 10))
	if err != nil {
		commentList, err = s.CommentMysqlManager.GetCommentList(req.VideoId)
		if err != nil {
			resp = &rpc_interact.CommentListResponse{
				BaseResp: &rpc_base.Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg:  "Error Get Comment List",
				},
			}
			return resp, err
		}
	}else {
		for _ , o := range commentIdListStr {
			tmp := new(model.Comment)
			exist , err := redis.GetJson(ctx , configs.SingleCommentHead + o , tmp)
			fmt.Println(tmp)
			if err != nil || !exist{
				resp = &rpc_interact.CommentListResponse{
					BaseResp: &rpc_base.Response{
						StatusCode: http.StatusInternalServerError,
						StatusMsg:  "Error Get Comment",
					},
				}
				return resp, err
			}
			commentList = append(commentList, tmp)
		}
	}
	for _, o := range commentList {
		timestamp := o.CreateDate
		seconds := timestamp / int64(time.Second)
		nanoseconds := timestamp % int64(time.Second)
		timeObj := time.Unix(seconds, nanoseconds)
		user, err := s.UserManager.GetUser(ctx, o.UserId)
		if err != nil {
			continue
		}
		timeStr := timeObj.Format("2006-01-02 15:04:05")
		resp.CommentList = append(resp.CommentList, &rpc_base.Comment{
			Id:         o.CommentID,
			User:       user,
			Content:    o.CommentText,
			CreateDate: timeStr,
		})
	}
	resp.BaseResp = &rpc_base.Response{
		StatusCode: http.StatusOK,
		StatusMsg:  "Successful Get Comment List",
	}
	return resp, nil
}
