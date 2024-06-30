package db

import "main/pkg/common"

// QueryPublishRecords 查询特定条件的上传记录 入参示例："user_id = ? and file_name = ?","{UserID}","{FileName}"
func QueryPublishRecords(query any, args ...any) ([]*common.VideoRecord, error) {
	res := []*common.VideoRecord{}
	err := db.Where(query, args).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, err
}
