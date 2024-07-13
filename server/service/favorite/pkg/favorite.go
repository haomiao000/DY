package pkg

import (
	"fmt"
	"main/server/common"
	_ "main/server/service/favorite/model"
	PublishModel "main/server/service/publish/model"
	"main/server/service/user/pkg"
	videoconf "main/server/service/video/model"
	videopkg "main/server/service/video/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	uid, exist := c.Get("uid")
	if !exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	if _, err := pkg.GetUser(uid.(int64)); err == nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	uid, exist := c.Get("uid")
	if !exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	if _, err := pkg.GetUser(uid.(int64)); err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	likeVideos, err := videopkg.QueryLikeVideos("uid = ? ", uid)
	if err != nil {
		fmt.Printf("query like videos error: %v", err)
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "query like video error"})
		return
	}
	videoList := []common.Video{}
	for _, likeVideo := range likeVideos {
		videoRecord := &videoconf.VideoRecord{}
		videoRecord, err = videopkg.GetVideoByID(likeVideo.VideoID)
		if err != nil {
			fmt.Printf("get video by id error: %v", err)
			continue
		}
		user, e := pkg.GetUser(videoRecord.UID)
		if e != nil {
			fmt.Printf("query like video id: %d author error: %v", videoRecord.UID, e)
			continue
		}
		videoList = append(videoList, common.Video{
			Id:            videoRecord.VideoID,
			Author:        *user,
			PlayUrl:       videoRecord.PlayUrl,
			CoverUrl:      videoRecord.CoverUrl,
			FavoriteCount: videoRecord.FavoriteCount,
			CommentCount:  videoRecord.CommentCount,
			IsFavorite:    true,
		})
	}
	c.JSON(http.StatusOK, PublishModel.VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
