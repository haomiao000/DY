package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"main/test/testcase"
	"main/pkg/common"
	"main/pkg/model"
	"time"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, model.FeedResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: testcase.DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
