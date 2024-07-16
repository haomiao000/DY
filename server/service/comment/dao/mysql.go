package dao
import (
	_"fmt"
	"main/internal/initialize"	
	"main/server/service/comment/model"
)
func CreateComment(comment *model.Comment) (error){
	err := initialize.DB.Create(comment).Error
	return err
}
func DeleteComment(commentID int64) (error) {
	err := initialize.DB.Delete(&model.Comment{}, commentID).Error
	return err
}
func GetComment(videoID int64) ([]*model.Comment , error) {
	var comments []*model.Comment
	err := initialize.DB.Where("video_id = ?" , videoID).Find(&comments).Error
	if err != nil {
		return nil , err
	}
	return comments , err
}