package pkg

import (
	"fmt"
	"github.com/haomiao000/DY/middleware"
	"github.com/haomiao000/DY/server/common"
	"github.com/haomiao000/DY/server/service/feed/model"
	videopkg "github.com/haomiao000/DY/server/service/video/pkg"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		token = string(c.PostForm("token"))
		if token != "" {
			claims, err := middleware.ParseToken(token)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "token is ERROR"})
				return
			}
			c.Set("userID", claims.UserID)
		}
	}
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
	// fmt.Printf("videolistlen: %d, videolist: %+v", len(videoList), videoList)
	c.JSON(http.StatusOK, model.FeedResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().UnixNano(),
	})
}
