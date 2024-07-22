package pkg

import (
	"main/server/service/relation/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"main/server/common"
	"main/server/service/relation/model"
)

func GetFollowMap(userID int64 , c *gin.Context) (*map[int64]bool) {
	var mp map[int64]bool
	if userList , err := dao.GetFollowUserList(userID); err != nil {
		c.JSON(http.StatusInternalServerError , model.UserListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg: "get follow user list error",
			},	
			UserList: nil,
		})
		return nil
	}else {
		mp = make(map[int64]bool)
		for _ , o := range userList {
			mp[o.UserID] = true
		}
	}
	return &mp
}