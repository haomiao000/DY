package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/haomiao000/DY/internal/grpc_gen/rpc_base"
	rpc_video "github.com/haomiao000/DY/internal/grpc_gen/rpc_video"
	client "github.com/haomiao000/DY/server/base_serv/video/api_client"
	model "github.com/haomiao000/DY/server/base_serv/video/model"
	"github.com/haomiao000/DY/server/common/configs"
	gorm "gorm.io/gorm"
)

type MysqlManager struct {
	videoDB *gorm.DB
}

func (m *MysqlManager) UpdateFavoriteCount(req *rpc_video.UpdateVideoFavoriteCountRequest) error {
	err := m.videoDB.Model(&model.VideoRecord{}).Where("video_id = ?", req.VideoId).
		UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", req.ChangeNumber)).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Printf("video is not exist in db,%s\n", err)
		return err
	}
	if err != nil {
		fmt.Println("update favorite count error")
		return err
	}
	return nil
}

func (m *MysqlManager) GetFavoriteVideoList(req *rpc_video.GetFavoriteVideoListByVideoIdRequest) ([]*model.VideoRecord, error) {
	var videoRecords []*model.VideoRecord
	err := m.videoDB.Where("video_id IN ?", req.VideoId).Find(&videoRecords).Error
	if err != nil {
		return nil, err
	}
	return videoRecords, nil
}

func (m *MysqlManager) GetAllVideo(ctx context.Context, req *rpc_video.GetFeedsReq) ([]*rpc_base.Video, error) {
	var videoRecords []*model.VideoRecord
	err := m.videoDB.Find(&videoRecords).Error
	if err != nil {
		return nil, err
	}
	var videos []*rpc_base.Video
	userMap, err := getUserByVideoRecord(ctx, videoRecords)
	if err != nil {
		return nil, err
	}
	for _, record := range videoRecords {
		user := userMap[record.UserID]
		videos = append(videos, &rpc_base.Video{
			Id: record.VideoID,
			Author: &rpc_base.User{
				Id:            user.GetId(),
				Name:          user.GetName(),
				FollowCount:   user.GetFollowCount(),
				FollowerCount: user.GetFollowerCount(),
				IsFollow:      user.GetIsFollow(),
			},
			PlayUrl:       record.PlayUrl,
			CoverUrl:      record.CoverUrl,
			FavoriteCount: record.FavoriteCount,
			CommentCount:  record.CommentCount,
			IsFavorite:    false,
		})
	}
	return videos, nil
}

func (m *MysqlManager) PublishVideo(req *rpc_video.PublishVideoReq) error {
	videoID, err := genVideoID()
	if err != nil {
		return err
	}
	// 保存视频数据
	return m.videoDB.Create(&model.VideoRecord{
		VideoID:       videoID,
		UserID:        req.UserId,
		FileName:      req.FileName,
		UpdateTime:    time.Now().Unix(),
		PlayUrl:       "",
		CoverUrl:      "",
		FavoriteCount: 0,
		CommentCount:  0,
	}).Error
}

func (m *MysqlManager) GetPublishVideo(ctx context.Context, userID int64) ([]*rpc_base.Video, error) {
	var videoRecords []*model.VideoRecord
	err := m.videoDB.Where("user_id = ?", userID).Find(&videoRecords).Error
	if err != nil {
		return nil, err
	}
	users, err := getUserByVideoRecord(ctx, videoRecords)
	var videos []*rpc_base.Video
	for _, record := range videoRecords {
		user := users[record.UserID]
		videos = append(videos, &rpc_base.Video{
			Id: record.VideoID,
			Author: &rpc_base.User{
				Id:            user.GetId(),
				Name:          user.GetName(),
				FollowCount:   user.GetFollowCount(),
				FollowerCount: user.GetFollowerCount(),
				IsFollow:      user.GetIsFollow(),
			},
			PlayUrl:       record.PlayUrl,
			CoverUrl:      record.CoverUrl,
			FavoriteCount: record.FavoriteCount,
			CommentCount:  record.CommentCount,
			IsFavorite:    false,
		})
	}
	return videos, nil
}

func (m *MysqlManager) UpdateCommentCount(req *rpc_video.UpdateVideoCommentCountRequest) error {
	err := m.videoDB.Model(&model.VideoRecord{}).Where("video_id = ?", req.VideoId).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", req.ChangeNumber)).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Printf("video is not exist in db,%s\n", err)
		return err
	}
	if err != nil {
		fmt.Println("update comment count error")
		return err
	}
	return nil
}

func NewMysqlManager(db *gorm.DB) *MysqlManager {
	m := db.Migrator()
	if !m.HasTable(&model.VideoRecord{}) {
		err := m.CreateTable(&model.VideoRecord{})
		if err != nil {
			fmt.Printf("create mysql table failed,%s\n", err)
		}
	}
	return &MysqlManager{videoDB: db}
}

func genVideoID() (int64, error) {
	node, err := snowflake.NewNode(configs.VideoSnowFlakeNode)
	if err != nil {
		return -1, err
	}
	return node.Generate().Int64(), nil
}

func getUserByVideoRecord(ctx context.Context, records []*model.VideoRecord) (map[int64]*rpc_base.User, error) {
	users := map[int64]bool{}
	var userList []int64
	for _, record := range records {
		users[record.UserID] = true
	}
	for userID := range users {
		userList = append(userList, userID)
	}
	userMap, err := client.GetUser(ctx, userList)
	if err != nil {
		return nil, err
	}
	return userMap, nil
}
