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
	favoriteStatus , err := dao.GetFavoriteStatus(userID.(int64), favoriteActionRequest.VideoID); 
	if err != nil {
		c.JSON(http.StatusNotFound, common.Response{StatusCode: 1 , StatusMsg: "get favorite status error"})
		return
	}
	if favoriteStatus == configs.IsLike {
		if favoriteActionRequest.ActionType == configs.Like {
			c.JSON(http.StatusOK, common.Response{StatusCode: 0 , StatusMsg: "you like the video you like"})
			return
		}else {
			if err := dao.DeleteFavorite(userID.(int64) , favoriteActionRequest.VideoID); err != nil {
				c.JSON(http.StatusInternalServerError , common.Response{StatusCode: 1, StatusMsg: "delete favorite error"})
				return
			}
			if err := dao.UpdateVideoFavoriteCound(favoriteActionRequest.VideoID , configs.Minus_like); err != nil {
				c.JSON(http.StatusInternalServerError , common.Response{StatusCode: 1, StatusMsg: "update video favorite count error"})
				return
			}
			c.JSON(http.StatusOK , common.Response{StatusCode: 0 , StatusMsg: "delete favrite successful"})
		}
	}else {
		if favoriteActionRequest.ActionType == configs.UnLike {
			c.JSON(http.StatusOK , common.Response{StatusCode: 0 , StatusMsg: "you unlike the video you unlike"})
			return
		}else {
			favorite := &model.Favorite{
				UserID: userID.(int64),
				VideoID: favoriteActionRequest.VideoID,
				CreateDate: time.Now().UnixNano(),
			}
			if err := dao.CreateFavorite(favorite); err != nil {
				c.JSON(http.StatusInternalServerError , common.Response{StatusCode: 1 , StatusMsg: "create favorite error"})
				return;
			}
			if err := dao.UpdateVideoFavoriteCound(favoriteActionRequest.VideoID , configs.Plus_like); err != nil {
				c.JSON(http.StatusInternalServerError , common.Response{StatusCode: 1, StatusMsg: "update video favorite count error"})
				return
			}
			c.JSON(http.StatusOK , common.Response{StatusCode: 0, StatusMsg: "create favorite successful"})
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
