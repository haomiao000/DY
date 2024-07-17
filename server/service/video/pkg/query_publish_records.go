package pkg

import (
	"main/internal/initialize"
	videoModel "main/server/service/video/model"
)

// QueryPublishRecords 查询特定条件的上传记录 入参示例WithUserID(1)
func QueryPublishRecords(args ...string) ([]*videoModel.VideoRecord, error) {
	query := assembleSQL(args...)
	tx := initialize.DB.Begin()
	res := []*videoModel.VideoRecord{}
	err := tx.Where(query).Find(&res).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return res, err
}
