package pkg

import (
	videoModel "github.com/haomiao000/DY/server/base_serv/video/model"
)

// GetAllVideo 从数据库中读取所有视频
func GetAllVideo() ([]*videoModel.VideoRecord, error) {
	return QueryPublishRecords()
}
