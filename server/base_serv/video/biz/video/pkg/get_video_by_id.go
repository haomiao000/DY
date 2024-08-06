package pkg

import (
	"fmt"
	"sync"

	videoModel "github.com/haomiao000/DY/server/base_serv/video/model"
)

var videoInfo sync.Map
var videoCount int64 = 0

// GetVideoByID 根据videoID查询视频
func GetVideoByID(videoID int64) (*videoModel.VideoRecord, error) {
	videoRecord := &videoModel.VideoRecord{}
	v, ok := videoInfo.Load(videoID)
	if !ok {
		// 未命中缓存，去数据库中找
		record, err := QueryPublishRecords(WithVideoID(videoID))
		if err != nil || len(record) != 1 {
			return nil, fmt.Errorf("video not exist, id: %d", videoID)
		}
		// 存到缓存中
		videoInfo.Store(record[0].VideoID, record[0])
		videoRecord = record[0]
	} else {
		videoRecord, ok = v.(*videoModel.VideoRecord)
		if !ok {
			return nil, fmt.Errorf("video not exist, id: %d", videoID)
		}
	}
	return videoRecord, nil
}
