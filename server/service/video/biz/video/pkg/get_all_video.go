package pkg

import (
	videoModel "main/server/service/video/model"
)

// GetAllVideo 从数据库中读取所有视频
func GetAllVideo() ([]*videoModel.VideoRecord, error) {
	return QueryPublishRecords()
}
