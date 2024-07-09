package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"main/test/testcase"
	"main/server/common"
	"main/server/service/feed/model"
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
