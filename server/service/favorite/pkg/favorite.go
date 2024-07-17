package pkg

import (
	"main/server/common"
	"main/server/service/favorite/dao"
	"main/server/service/favorite/model"
	PublishModel "main/server/service/publish/model"

	"main/configs"
	"main/test/testcase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	var favoriteActionRequest model.FavoriteActionRequest
	if err := c.ShouldBind(&favoriteActionRequest); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "missing field"})
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not logged in"})
		return
	}
	if favoriteType, err := dao.GetFavoriteType(userID.(int64), favoriteActionRequest.VideoID); err != nil {
		if favoriteActionRequest.ActionType == configs.UnLike {
			c.JSON(http.StatusOK, gin.H{"message": "before unlike , without like"})
			return
		} else if favoriteActionRequest.ActionType == configs.Like {
			if err := dao.CreateFavorite(userID.(int64), favoriteActionRequest.VideoID, favoriteActionRequest.ActionType); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "create favorite error"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "create favorite successful"})
			return
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "error actiontype"})
			return
		}
	} else {
		if favoriteActionRequest.ActionType == favoriteType {
			c.JSON(http.StatusOK, gin.H{"message": "same action type"})
			return
		} else if favoriteActionRequest.ActionType == configs.Like || favoriteActionRequest.ActionType == configs.UnLike {
			if err := dao.UpdateFavoriteActionType(userID.(int64), favoriteActionRequest.VideoID, favoriteActionRequest.ActionType); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "update actionType error"})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "favorite update successful"})
				return
			}
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "error favorite actionType"})
		}
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
