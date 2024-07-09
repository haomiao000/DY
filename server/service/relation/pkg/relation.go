package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"main/server/common"
	"main/test/testcase"
	"main/server/service/relation/model"
	"main/server/service/user/pkg"
)


// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := pkg.UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, model.UserListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		UserList: []common.User{testcase.DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, model.UserListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		UserList: []common.User{testcase.DemoUser},
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, model.UserListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		UserList: []common.User{testcase.DemoUser},
	})
}
