package dao

import (
	model "github.com/haomiao000/DY/server/service/relation/model"
	rpc_relation "github.com/haomiao000/DY/server/grpc_gen/rpc_relation"
	gorm "gorm.io/gorm"
	"fmt"
)
type MysqlManager struct {
	concernDB *gorm.DB
}

func (m *MysqlManager) GetRelationStatus(req *rpc_relation.RelationActionRequest) (bool, error) {
	var concern model.ConcernsInfo
	err := m.concernDB.Where("user_id = ? AND follower_id = ?" , req.ToUserId, req.UserId).First(&concern).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil { 
		fmt.Println("select relation status error")
		return false, err
	}else {
		return true, nil
	}
}

func (m *MysqlManager) CreateRelationInfo(userID int64 , toUserID int64) (error) {
	err := m.concernDB.Create(&model.ConcernsInfo{
		UserID: toUserID,
		FollowerID: userID,
	}).Error
	if err != nil {
		fmt.Printf("create relation between %d and %d error\n" , userID , toUserID)
		return err
	}
	return nil
}

func (m *MysqlManager) DeleteRelationInfo(userID int64 , toUserID int64) (error) {
	err := m.concernDB.Where("user_id = ? AND follower_id = ?" , toUserID , userID).Delete(&model.ConcernsInfo{}).Error
	if err != nil {
		fmt.Println("delete concern error")
		return err
	}
	return nil
}

func (m *MysqlManager) GetFollowUserIdList(userID int64) ([]int64 , error) {
	var userIDs []int64
	err := m.concernDB.Model(&model.ConcernsInfo{}).
		Where("follower_id = ?", userID).
		Pluck("user_id", &userIDs).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("user follow user is empty")
		return nil , nil
	}
	if err != nil {
		return nil, err
	}
	return userIDs, nil
}

func (m *MysqlManager) GetFollowerUserIdList(userID int64) ([]int64 , error) {
	var userIDs []int64
	err := m.concernDB.Model(&model.ConcernsInfo{}).
		Where("user_id = ?", userID).
		Pluck("follower_id", &userIDs).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("user follower user is empty")
		return nil , nil
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return userIDs, nil
}

func (m *MysqlManager) GetMutualFollowersIdList(userID int64) ([]int64 , error) {
	var mutualFollowers []*model.ConcernsInfo
    err := m.concernDB.Table("concerns_info as ci1").
        Select("ci1.user_id, ci1.follower_id").
        Joins("JOIN concerns_info as ci2 ON ci1.user_id = ci2.follower_id AND ci1.follower_id = ci2.user_id").
        Where("ci1.user_id = ? OR ci1.follower_id = ?", userID, userID).
        Scan(&mutualFollowers).Error

	if err == gorm.ErrRecordNotFound {
		fmt.Println("user friend user is empty")
		return nil , nil
	}
    if err != nil {
        return nil, err
    }
    var userIDs []int64
    for _, follower := range mutualFollowers {
        if follower.UserID != userID {
            userIDs = append(userIDs, follower.UserID)
        }
    }
	return userIDs , nil
}

func NewMysqlManager(db *gorm.DB) *MysqlManager {
	m := db.Migrator()
	if !m.HasTable(&model.ConcernsInfo{}) {
		err := m.CreateTable(&model.ConcernsInfo{})
		if err != nil {
			fmt.Printf("create mysql table failed,%s\n", err)
		}
	}
	return &MysqlManager{concernDB: db}
}
