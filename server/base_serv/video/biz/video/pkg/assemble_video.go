package pkg

import (
	// "fmt"
	userpkg "github.com/haomiao000/DY/server/base_serv/user/pkg"
	videoModel "github.com/haomiao000/DY/server/base_serv/video/model"
	common "github.com/haomiao000/DY/server/common"
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
			// fmt.Printf("get user info error: %v, userid: %v", e, record.UserID)
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
	// fmt.Printf("res: %+v", res)
	return res, nil
}
