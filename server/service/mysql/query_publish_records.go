package db

import "main/server/common"

// QueryPublishRecords 查询特定条件的上传记录 入参示例："user_id = ? and file_name = ?","{UserID}","{FileName}"
func QueryPublishRecords(query any, args ...any) ([]*common.VideoRecord, error) {
	tx := db.Begin()
	res := []*common.VideoRecord{}
	err := tx.Where(query, args).Find(&res).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return res, err
}
