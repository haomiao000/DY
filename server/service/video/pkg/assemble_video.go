package pkg

import (
	"fmt"
	"main/server/common"
	userpkg "main/server/service/user/pkg"
	videoModel "main/server/service/video/model"
)

func AssembleVideo(userID int64, videoRecord []*videoModel.VideoRecord) (map[int64]*common.Video, error) {
	likeVideosID, err := QueryLikeVideos(WithUserID(userID))
	if err != nil {
		return nil, err
	}
	like := map[int64]bool{}
	for _, id := range likeVideosID {
		like[id] = true
	}
	res := map[int64]*common.Video{}
	for _, record := range videoRecord {
		user, e := userpkg.GetUser(record.UserID)
		if e != nil {
			fmt.Printf("get user info error: %v, userid: %v", e, record.UserID)
			continue
		}
		res[record.VideoID] = &common.Video{
			Id:            record.VideoID,
			Author:        *user,
			PlayUrl:       record.PlayUrl,
			CoverUrl:      record.CoverUrl,
			FavoriteCount: record.FavoriteCount,
			CommentCount:  record.CommentCount,
			IsFavorite:    like[record.VideoID],
		}
	}
	return res, nil
}
