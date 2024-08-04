package dao

import (
	rpc_interact "github.com/haomiao000/DY/server/grpc_gen/rpc_interact"
	model "github.com/haomiao000/DY/server/service/interact/model"
	gorm "gorm.io/gorm"
	"fmt"
)
type MysqlManager struct {
	commentDB *gorm.DB 
	favoriteDB *gorm.DB
}

func (m *MysqlManager) GetFavoriteStatus(req *rpc_interact.FavoriteActionRequest) (bool, error){
	var favorite model.Favorite
	err := m.favoriteDB.Where("user_id = ? AND video_id = ?" , req.UserId, req.VideoId).First(&favorite).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil { 
		fmt.Println("select favorite status error")
		return false, err
	}else {
		return true, nil
	}
}

func (m *MysqlManager) DeleteFavorite(req *rpc_interact.FavoriteActionRequest) (error) {
	err := m.favoriteDB.Where("user_id = ? AND video_id = ?" , req.UserId , req.VideoId).Delete(&model.Favorite{}).Error
	if err != nil {
		fmt.Println("delete favorite error")
		return err
	}
	return nil
}

func (m *MysqlManager) CreateFavorite(favorite *model.Favorite) (error) {
	err := m.favoriteDB.Create(favorite).Error
	if err != nil {
		fmt.Println("create favorite error")
		return err
	}
	return nil
}

func (m *MysqlManager) GetFavoriteVideoIdList(req *rpc_interact.FavoriteListRequest) ([]int64 , error) {
	var favorites []model.Favorite
	err := m.favoriteDB.Where("user_id = ?", req.OwnerId).Find(&favorites).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("user favorite video is empty")
		return nil , nil
	}
	if err != nil {
		fmt.Println("get favorite video id list error")
		return nil , err
	}
	videoIDs := make([]int64, len(favorites))
	for i, favorite := range favorites {
		videoIDs[i] = favorite.VideoID
	}
	return videoIDs , nil
}

func (m *MysqlManager) CreateComment(comment *model.Comment) (error) {
	err := m.commentDB.Create(comment).Error
	if err != nil {
		fmt.Println("create comment error")
		return err
	}
	return nil
}

func (m *MysqlManager) DeleteComment(commentID int64) (error) {
	err := m.commentDB.Delete(&model.Comment{}, commentID).Error
	if err != nil {
		fmt.Println("delete comment error")
		return err
	}
	return nil
}

func (m *MysqlManager) GetCommentList(videoID int64) ([]*model.Comment , error) {
	var comments []*model.Comment
	err := m.commentDB.Where("video_id = ?" , videoID).Find(&comments).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("this video comment is empty")
		return nil , nil
	}
	if err != nil {
		return nil , err
	}
	return comments , err
}


func NewMysqlManager(db *gorm.DB) *MysqlManager {
	m := db.Migrator()
	if !m.HasTable(&model.Favorite{}) {
		err := m.CreateTable(&model.Favorite{})
		if err != nil {
			fmt.Printf("create mysql table failed,%s\n", err)
		}
	}
	if !m.HasTable(&model.Comment{}) {
		err := m.CreateTable(&model.Comment{})
		if err != nil {
			fmt.Printf("create mysql table failed,%s\n", err)
		}
	}
	return &MysqlManager{
		commentDB: db,
		favoriteDB: db,
	}
}
