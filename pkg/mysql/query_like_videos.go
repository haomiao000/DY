package db

import "main/pkg/common"

// QueryLikeVideos 查询特定条件的点赞视频 入参示例："user_id = ?","{UserID}"
func QueryLikeVideos(query any, args ...any) ([]*common.LikeVideo, error) {
	tx := db.Begin()
	res := []*common.LikeVideo{}
	err := tx.Where(query, args).Find(&res).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return res, err
}
