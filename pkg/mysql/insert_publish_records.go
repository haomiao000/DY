package db

import (
	"main/pkg/common"
)

// InsertPublishRecords 插入记录
func InsertPublishRecords(records []*common.VideoRecord) error {
	tx := db.Begin()
	for _, record := range records {
		err := tx.Create(record).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return nil
}
