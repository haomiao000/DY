package pkg

import (
	"github.com/haomiao000/DY/internal/initialize"
	favoriteModel "github.com/haomiao000/DY/server/base_serv/favorite/model"

	"gorm.io/gorm"
)

// QueryLikeVideos 查询特定条件的点赞视频 入参示例withVideoID(1)
func QueryLikeVideos(args ...string) ([]int64, error) {
	query := assembleSQL(args...)
	tx := initialize.DB.Begin()
	favorite := []*favoriteModel.Favorite{}
	err := tx.Where(query).Find(&favorite).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	res := []int64{}
	for _, v := range favorite {
		res = append(res, v.VideoID)
	}
	return res, nil
}
