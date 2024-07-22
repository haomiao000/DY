package dao

import (
	"errors"
	_ "fmt"
	"main/internal/initialize"
	"main/server/service/relation/model"
	UserModel "main/server/service/user/model"

	"gorm.io/gorm"
)

func CreateRelationInfo(userID int64 , toUserID int64) (error) {
	err := initialize.DB.Create(&model.ConcernsInfo{
		UserID: toUserID,
		FollowerID: userID,
	}).Error
	return err
}
func FindRelationInfo(userID int64, toUserID int64) (bool, error) {
	var concern model.ConcernsInfo
	err := initialize.DB.Where("user_id = ? AND follower_id = ?" , toUserID, userID).First(&concern).Error
	
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil { 
		return false, err
	}

	return true, nil
}
func DeleteRelationInfo(userID int64 , toUserID int64) (error) {
	err := initialize.DB.Where("user_id = ? AND follower_id = ?" , toUserID , userID).Delete(&model.ConcernsInfo{}).Error
	return err
}
//查询关注列表
func GetFollowUserList(ownerID int64) ([]*UserModel.User , error) {
	var userList []*UserModel.User
	err := initialize.DB.
	Table("concerns_info").
	Joins("JOIN user AS u ON concerns_info.user_id = u.user_id").
	Where("concerns_info.follower_id = ?", ownerID).
	Select("u.user_id, u.name, u.followcount, u.followercount").
	Find(&userList).
	Error
	return userList , err 
}
//查询粉丝列表
func GetFollowerUserList(ownerID int64) ([]*UserModel.User , error) {
	var userList []*UserModel.User
	err := initialize.DB.
	Table("concerns_info").
	Joins("JOIN user AS u ON concerns_info.follower_id = u.user_id").
	Where("concerns_info.user_id = ?", ownerID).
	Select("u.user_id, u.name, u.followcount, u.followercount").
	Find(&userList).
	Error
	return  userList , err
}
func GetMutualFollowers(userID int64) ([]*UserModel.User, error) {
    var mutualFollowers []*model.ConcernsInfo

    err := initialize.DB.Table("concerns_info as ci1").
        Select("ci1.user_id, ci1.follower_id").
        Joins("JOIN concerns_info as ci2 ON ci1.user_id = ci2.follower_id AND ci1.follower_id = ci2.user_id").
        Where("ci1.user_id = ? OR ci1.follower_id = ?", userID, userID).
        Scan(&mutualFollowers).Error

    if err != nil {
        return nil, err
    }
    var userIDs []int64
    for _, follower := range mutualFollowers {
        if follower.UserID != userID {
            userIDs = append(userIDs, follower.UserID)
        }
    }
    var users []*UserModel.User
    if len(userIDs) > 0 {
        err = initialize.DB.Where("user_id IN ?", userIDs).Find(&users).Error
        if err != nil {
            return nil, err
        }
    }
    return users, nil
}