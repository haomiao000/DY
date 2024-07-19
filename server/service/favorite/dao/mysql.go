package dao

import (
	_ "fmt"
	"errors"
	"gorm.io/gorm"
	"main/internal/initialize"
	videoModel "main/server/service/video/model"
	"main/server/service/favorite/model"
)

func GetFavoriteStatus(userID int64, videoID int64) (bool, error) {
	var favorite model.Favorite
	err := initialize.DB.Where("user_id = ? AND video_id = ?" , userID, videoID).First(&favorite).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil { 
		return false, err
	}
	return true, nil
}
func CreateFavorite(favorite *model.Favorite) (error) {
	err := initialize.DB.Create(favorite).Error
	return err
}

func DeleteFavorite(userID int64 , videoID int64) (error) {
	err := initialize.DB.Where("user_id = ? AND video_id = ?" , userID , videoID).Delete(&model.Favorite{}).Error
	return err
}
func GetFavoriteList(userID int64) ([]*model.Favorite , error) {
	var favorites []*model.Favorite
	err := initialize.DB.Where("user_id = ?" , userID).Find(&favorites).Error
	if err != nil {
		return nil , err
	}
	return favorites , err
}
func GetFavoriteVideoListByUserID(userID int64) ([]*videoModel.VideoRecord , error) {
	var videos []*videoModel.VideoRecord
	err := initialize.DB.
			Joins("JOIN favorite ON video_records.video_id = favorite.video_id").
			Where("favorite.user_id = ?", userID).
			Find(&videos).Error;
	if err != nil {
		return nil , err
	}
	return videos , err
}