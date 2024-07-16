package dao
import (
	_"fmt"
	"main/internal/initialize"	
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
