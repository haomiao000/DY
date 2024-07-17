package pkg

import (
	"fmt"
	"main/server/common"
	"main/server/service/favorite/dao"
	"main/server/service/favorite/model"
	PublishModel "main/server/service/publish/model"

	"main/configs"
	// "main/test/testcase"
	UserService "main/server/service/user/pkg"
	"net/http"
	"time"

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
		}else if favoriteActionRequest.ActionType == configs.Like {
			favorite := &model.Favorite{
				UserID: userID.(int64),
				VideoID: favoriteActionRequest.VideoID,
				ActionType: favoriteActionRequest.ActionType,
				CreateDate: time.Now().UnixNano(),
			}
			if err := dao.CreateFavorite(favorite);err != nil {
				c.JSON(http.StatusInternalServerError , gin.H{"error" : "create favorite error"})
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
	userID , exists := c.Get("userID") 
	if !exists {
		c.JSON(http.StatusUnauthorized , gin.H{"error" : "user not logged in"})
		return
	}
	favoriteVideoList , err := dao.GetFavoriteVideoListByUserID(userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError , gin.H{"error" : "Get favorite list error"})
		return
	}
	var videoListResponse PublishModel.VideoListResponse
	for _ , o := range favoriteVideoList {
		auth , err :=  UserService.GetUser(o.UserID)
		if err != nil {
			fmt.Println("GetUser  error")
			continue
		}
		videoListResponse.VideoList = append(videoListResponse.VideoList, &common.Video{
			Id: o.VideoID,
			Author: *auth,
			PlayUrl: o.PlayUrl,
			CoverUrl: o.CoverUrl,
			FavoriteCount: o.FavoriteCount,
			CommentCount: o.CommentCount,
			IsFavorite: true,
		})
	}
	videoListResponse.BaseResp = &common.Response{
		StatusCode: http.StatusOK,
		StatusMsg: "get favorite list successful",
	}
	c.JSON(http.StatusOK , videoListResponse)
}
