package db

import "main/pkg/common"

// QueryLikeVideos 查询特定条件的点赞视频 入参示例："user_id = ?","{UserID}"
func QueryLikeVideos(query any, args ...any) ([]*common.LikeVideo, error) {
	res := []*common.LikeVideo{}
	err := db.Where(query, args).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, err
}
