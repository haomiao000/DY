package pkg

import (
	"main/configs"
	"main/internal/initialize"
	favoritemodel "main/server/service/favorite/model"
)

// QueryLikeVideos 查询特定条件的点赞视频 入参示例withVideoID(1)
func QueryLikeVideos(args ...string) ([]int64, error) {
	query := assembleSQL(args...)
	tx := initialize.DB.Begin()
	favorite := []*favoritemodel.Favorite{}
	err := tx.Where(query).Find(&favorite).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	res := []int64{}
	for _, v := range favorite {
		if v.ActionType == configs.IsLike {
			res = append(res, v.VideoID)
		}
	}
	return res, err
}
