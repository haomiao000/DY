package pkg

import (
	"github.com/gin-gonic/gin"
	"main/server/common"
	"main/test/testcase"
	"main/server/service/comment/model"
	"net/http"
	"main/server/service/user/pkg"
)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	if user, exist := pkg.UsersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			c.JSON(http.StatusOK, model.CommentActionResponse{Response: common.Response{StatusCode: 0},
				Comment: common.Comment{
					Id:         1,
					User:       user,
					Content:    text,
					CreateDate: "05-01",
				}})
			return
		}
		c.JSON(http.StatusOK, common.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, model.CommentListResponse{
		Response:    common.Response{StatusCode: 0},
		CommentList: testcase.DemoComments,
	})
}
