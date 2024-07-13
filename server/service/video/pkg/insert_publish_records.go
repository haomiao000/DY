package pkg

import (
	"main/internal/initialize"
	videoconf "main/server/service/video/model"
	"sync/atomic"
)

// InsertPublishRecords 插入记录 videoID不需要填，在函数内会填上 同时会将数据保存到缓存中
func InsertPublishRecords(records []*videoconf.VideoRecord) error {
	tx := initialize.DB.Begin()
	for _, record := range records {
		// 间隙锁的原因，videoID一定连续
		record.VideoID = atomic.AddInt64(&videoCount, 1)
		err := tx.Create(record).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	// 事务提交后再存入缓存，避免事务回滚时数据已经存入到缓存中
	for _, record := range records {
		videoInfo.Store(record.VideoID, record)
	}
	return nil
}
