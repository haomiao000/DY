package pkg

import (
	"github.com/haomiao000/DY/internal/initialize"
	videoModel "github.com/haomiao000/DY/server/service/video/model"

	"gorm.io/gorm"
)

// QueryPublishRecords 查询特定条件的上传记录 入参示例WithUserID(1)
func QueryPublishRecords(args ...string) ([]*videoModel.VideoRecord, error) {
	query := assembleSQL(args...)
	tx := initialize.DB.Begin()
	res := []*videoModel.VideoRecord{}
	err := tx.Where(query).Find(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return res, err
}
