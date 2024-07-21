package pkg

import (
	"fmt"
	"main/server/common"
	"main/server/service/feed/model"
	videopkg "main/server/service/video/pkg"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// 展示所有作品
	// fmt.Println("----------")
	videoRecords, err := videopkg.GetAllVideo()
	if err != nil {
		fmt.Printf("get video error: %v", err)
		c.JSON(http.StatusOK, model.FeedResponse{
			Response:  common.Response{StatusCode: 1},
			VideoList: []*common.Video{},
			NextTime:  time.Now().Unix(),
		})
		return
	}
	// 装填videos
	v, exist := c.Get("userID")
	userID := int64(0)
	if exist {
		userID = v.(int64)
	}
	videos, err := videopkg.AssembleVideo(userID, videoRecords)
	if err != nil {
		fmt.Printf("assemble videos error: %v, videorecords: %+v", err, videoRecords)
		c.JSON(http.StatusOK, model.FeedResponse{
			Response:  common.Response{StatusCode: 1},
			VideoList: []*common.Video{},
			NextTime:  time.Now().Unix(),
		})
		return
	}

	videoList := []*common.Video{}
	for _, v := range videos {
		videoList = append(videoList, v)
	}
	fmt.Printf("videolistlen: %d, videolist: %+v", len(videoList), videoList)

	// fmt.Println("-----")
	// fmt.Println("-----")
	c.JSON(http.StatusOK, model.FeedResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().UnixNano(),
	})
}
