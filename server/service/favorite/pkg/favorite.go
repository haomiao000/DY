package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"main/test/testcase"
	"main/server/common"
	_ "main/server/service/favorite/model"
	PublishModel "main/server/service/publish/model"
	"main/server/service/user/pkg"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := pkg.UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, PublishModel.VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: testcase.DemoVideos,
	})
}
