package dao

import (
	"fmt"

	rpc_video "github.com/haomiao000/DY/internal/grpc_gen/rpc_video"
	model "github.com/haomiao000/DY/server/base_serv/video/model"
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
