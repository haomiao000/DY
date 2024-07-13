package pkg

import (
	"main/internal/initialize"
	videoconf "main/server/service/video/model"
)

// QueryPublishRecords 查询特定条件的上传记录 入参示例："uid = ? and file_name = ?","{UserID}","{FileName}"
func QueryPublishRecords(query any, args ...any) ([]*videoconf.VideoRecord, error) {
	tx := initialize.DB.Begin()
	res := []*videoconf.VideoRecord{}
	err := tx.Where(query, args).Find(&res).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return res, err
}
