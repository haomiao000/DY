package pkg

import (
	"fmt"
	videoconf "main/server/service/video/model"
	"sync"
)

var videoInfo sync.Map
var videoCount int64 = 0

// GetVideoByID 根据videoID查询视频
func GetVideoByID(videoID int64) (*videoconf.VideoRecord, error) {
	videoRecord := &videoconf.VideoRecord{}
	v, ok := videoInfo.Load(videoID)
	if !ok {
		// 未命中缓存，去数据库中找
		record, err := QueryPublishRecords("video_id = ?", videoID)
		if err != nil || len(record) != 1 {
			return nil, fmt.Errorf("video not exist, id: %d", videoID)
		}
		// 存到缓存中
		videoInfo.Store(record[0].VideoID, record[0])

	} else {
		videoRecord, ok = v.(*videoconf.VideoRecord)
		if !ok {
			return nil, fmt.Errorf("video not exist, id: %d", videoID)
		}
	}
	return videoRecord, nil
}
