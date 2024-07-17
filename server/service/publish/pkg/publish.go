package pkg

import (
	"fmt"
	"main/configs"
	"main/server/common"
	"main/server/service/publish/model"
	videoModel "main/server/service/video/model"
	videopkg "main/server/service/video/pkg"
	"net/http"
	"path/filepath"
	"time"

	"main/server/service/user/pkg"

	"github.com/gin-gonic/gin"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	user, err := pkg.GetUser(userID.(int64))
	if !exist {
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
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./assets/public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	video := &videoModel.VideoRecord{ // TODO 记录信息不完整，待补充
		UserID:        user.Id,
		FileName:      finalName,
		UpdateTime:    time.Now().UnixMilli(),
		PlayUrl:       configs.VideoURL + finalName,
		CoverUrl:      configs.VideoURL + finalName,
		FavoriteCount: 0,
		CommentCount:  0,
	}
	if err = videopkg.InsertPublishRecords([]*videoModel.VideoRecord{video}); err != nil {
		fmt.Printf("%v uploaded error: %v", finalName, err)
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 0,
			StatusMsg:  finalName + " uploaded failed",
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	user, err := pkg.GetUser(userID.(int64))
	if !exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	videos, err := videopkg.QueryPublishRecords(videopkg.WithUserID(user.Id))
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "get video error"})
		return
	}
	// 查询用户点赞过的视频
	likeVideosList, err := videopkg.QueryLikeVideos(videopkg.WithUserID(user.Id))
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "get like video error"})
		return
	}
	// 用户点赞过的视频记录下来
	likeVideos := make(map[int64]bool, len(likeVideosList))
	for _, likeVideo := range likeVideosList {
		likeVideos[likeVideo] = true
	}
	videoList := []*common.Video{}
	for _, video := range videos {
		videoRecord := &videoModel.VideoRecord{}
		videoRecord, err = videopkg.GetVideoByID(video.VideoID)
		if err != nil {
			fmt.Printf("get video by id error: %v", err)
			return
		}
		videoList = append(videoList, &common.Video{
			Id:            videoRecord.VideoID,
			Author:        *user,
			PlayUrl:       videoRecord.PlayUrl,
			CoverUrl:      videoRecord.CoverUrl,
			FavoriteCount: videoRecord.FavoriteCount,
			CommentCount:  videoRecord.CommentCount,
			IsFavorite:    likeVideos[video.VideoID],
		})
	}

	c.JSON(http.StatusOK, model.VideoListResponse{
		BaseResp: &common.Response{
			StatusCode: 0,
		},
		VideoList: videoList, // TODO 补充逻辑后返回videoList
	})
}
