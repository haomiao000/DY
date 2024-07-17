package pkg

import (
	"fmt"
	"main/server/common"
	"main/server/service/feed/model"
	userpkg "main/server/service/user/pkg"
	videopkg "main/server/service/video/pkg"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// 展示videoID为1的作品
	// fmt.Println("----------")
	videoRecord, err := videopkg.GetVideoByID(1)
	if err != nil {
		fmt.Printf("get video by id error: %v", err)
		c.JSON(http.StatusOK, model.FeedResponse{
			Response:  common.Response{StatusCode: 1},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
		return
	}
	// 查作者信息
	user, err := userpkg.GetUser(videoRecord.UserID)
	if err != nil {
		fmt.Printf("get user error: %v", err)
		c.JSON(http.StatusOK, model.FeedResponse{
			Response:  common.Response{StatusCode: 1},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
		return
	}
	// check一下是否喜欢
	videoID, err := videopkg.QueryLikeVideos(videopkg.WithUserID(videoRecord.UserID),
		videopkg.WithVideoID(videoRecord.VideoID))
	if err != nil {
		fmt.Printf("query favorite info error: %v", err)
		c.JSON(http.StatusOK, model.FeedResponse{
			Response:  common.Response{StatusCode: 1},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
		return
	}
	isLike := false
	if len(videoID) == 1 {
		isLike = true
	}
	// fmt.Println("-----")
	fmt.Println(videoRecord.PlayUrl)
	// fmt.Println("-----")
	c.JSON(http.StatusOK, model.FeedResponse{
		Response: common.Response{StatusCode: 0},
		VideoList: []*common.Video{
			{
				Id:            videoRecord.VideoID,
				Author:        *user,
				PlayUrl:       videoRecord.PlayUrl,
				CoverUrl:      videoRecord.CoverUrl,
				FavoriteCount: videoRecord.FavoriteCount,
				CommentCount:  videoRecord.CommentCount,
				IsFavorite:    isLike,
			},
		},
		NextTime: time.Now().UnixNano(),
	})
}