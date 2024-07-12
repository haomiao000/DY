package pkg

import (
	"main/internal/initialize"
	videoconf "main/server/service/video/model"
)

// QueryLikeVideos 查询特定条件的点赞视频 入参示例："user_id = ?","{UserID}"
func QueryLikeVideos(query any, args ...any) ([]*videoconf.LikeVideo, error) {
	tx := initialize.DB.Begin()
	res := []*videoconf.LikeVideo{}
	err := tx.Where(query, args).Find(&res).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return res, err
}
