package dao

import (
	_ "fmt"
	"main/configs"
	"main/internal/initialize"
	"main/server/common"
	"main/server/service/favorite/model"
)

func GetFavoriteType(userID int64 , videoID int64) (int8 , error) {
	var favorite model.Favorite
	err := initialize.DB.Where("user_id = ? AND video_id = ?" , userID , videoID).First(&favorite).Error; 
	return favorite.ActionType , err
}
func CreateFavorite(favorite *model.Favorite) (error) {
	err := initialize.DB.Create(favorite).Error
	return err
}
func UpdateFavoriteActionType(userID int64 , videoID int64 , actionType int8) (error) {
	err := initialize.DB.Model(&model.Favorite{}).Where("user_id = ? AND video_id = ?" , userID , videoID).Update("action_type" , actionType).Error
	return err
}
func GetFavoriteList(userID int64) ([]*model.Favorite , error) {
	var favorites []*model.Favorite
	err := initialize.DB.Where("user_id = ? AND action_type = ?" , userID , configs.Like).Find(&favorites).Error
	if err != nil {
		return nil , err
	}
	return favorites , err
}
func GetFavoriteVideoListByUserID(userID int64) ([]*common.VideoRecord , error) {
	var videos []*common.VideoRecord
	err := initialize.DB.
		Joins("JOIN favorite ON video_records.video_id = favorite.video_id").
		Where("favorite.user_id = ? AND favorite.action_type = ?", userID, configs.Like).
		Find(&videos).Error;
	if err != nil {
		return nil , err
	}
	return videos , err
}