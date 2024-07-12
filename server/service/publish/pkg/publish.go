package pkg

import (
	"fmt"
	"main/server/common"
	"main/server/service/publish/model"
	videoconf "main/server/service/video/model"
	videopkg "main/server/service/video/pkg"
	"net/http"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"main/server/service/user/pkg"

	"github.com/gin-gonic/gin"
)

var videoInfo sync.Map
var videoCount int64 = 0

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := pkg.UsersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := pkg.UsersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./assets/public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	video := &videoconf.VideoRecord{ // TODO 记录信息不完整，待补充
		VideoID:       atomic.AddInt64(&videoCount, 1),
		UserID:        user.Id,
		FileName:      finalName,
		UpdateTime:    time.Now().UnixMilli(),
		PlayUrl:       "",
		CoverUrl:      "",
		FavoriteCount: 0,
		CommentCount:  0,
	}
	if err = videopkg.InsertPublishRecords([]*videoconf.VideoRecord{video}); err != nil {
		fmt.Printf("%v uploaded error: %v", finalName, err)
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 0,
			StatusMsg:  finalName + " uploaded failed",
		})
		return
	}
	videoInfo.Store(finalName, video)

	c.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	user, exist := pkg.UsersLoginInfo[token]
	if !exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	videos, err := videopkg.QueryPublishRecords("user_id = ?", user.Id)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "get video error"})
		return
	}
	// 查询用户点赞过的视频
	likeVideosList, err := videopkg.QueryLikeVideos("user_id = ?", user.Id)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "get like video error"})
		return
	}
	// 用户点赞过的视频记录下来
	likeVideos := make(map[int64]bool, len(likeVideosList))
	for _, likeVideo := range likeVideosList {
		likeVideos[likeVideo.VideoID] = true
	}
	videoList := []common.Video{}
	for _, video := range videos {
		v, ok := videoInfo.Load(video.FileName)
		if !ok {
			fmt.Printf("video not exist, video id: %d", video.VideoID)
			continue
		}
		video, ok := v.(*videoconf.VideoRecord)
		if !ok {
			continue
		}
		videoList = append(videoList, common.Video{
			Id:            video.VideoID,
			Author:        user,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    likeVideos[video.VideoID],
		})
	}

	c.JSON(http.StatusOK, model.VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: videoList, // TODO 补充逻辑后返回videoList
	})
}
