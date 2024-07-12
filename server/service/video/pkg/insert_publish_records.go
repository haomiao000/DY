package pkg

import (
	"main/internal/initialize"
	videoconf "main/server/service/video/model"
)

// InsertPublishRecords 插入记录
func InsertPublishRecords(records []*videoconf.VideoRecord) error {
	tx := initialize.DB.Begin()
	for _, record := range records {
		err := tx.Create(record).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
