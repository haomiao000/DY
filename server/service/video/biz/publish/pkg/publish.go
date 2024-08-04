package pkg

import (
	"fmt"
	"github.com/haomiao000/DY/configs"
	"github.com/haomiao000/DY/server/common"
	"main/server/service/publish/model"
	videoModel "github.com/haomiao000/DY/server/service/video/model"
	videopkg "github.com/haomiao000/DY/server/service/video/pkg"
	"net/http"
	"path/filepath"
	"time"

	"github.com/haomiao000/DY/server/service/user/pkg"

	"github.com/gin-gonic/gin"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	// fmt.Println("----------------")
	userID, exist := c.Get("userID")
	if !exist {
		fmt.Println("get userID error")
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	// fmt.Println("error 1")
	user, err := pkg.GetUser(userID.(int64))
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	// fmt.Println("error 1")
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	// fmt.Println("error 1")
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
	// fmt.Println("error 1")
	video := &videoModel.VideoRecord{
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
	// fmt.Println("error 1")
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
	if err != nil {
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
			// fmt.Printf("get video by id error: %v", err)
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
