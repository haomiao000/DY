package db

import (
	"fmt"
	"main/pkg/common"
)

// InsertPublishRecords 插入记录
func InsertPublishRecords(records []*common.VideoRecord) {
	for _, record := range records {
		err := db.Create(record).Error
		if err != nil {
			fmt.Printf("insert query: %+v error: %v", record, err)
		}
	}
}
