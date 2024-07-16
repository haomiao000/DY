package pkg

import (
	"main/configs"
	"main/server/common"
	"main/server/service/comment/dao"
	"main/server/service/comment/model"
	// "main/test/testcase"
	"net/http"
	"time"
	"strconv"
	UserServer "main/server/service/user/pkg"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	var commentActionRequest model.CommentActionRequest
	var commentActionResponse model.CommentActionResponse
	if err := c.ShouldBind(&commentActionRequest);err != nil {
		c.JSON(http.StatusNotFound , gin.H{"error" : "comment action bind error"})
		return
	}
	userID , exists := c.Get("userID") 
	if !exists {
		c.JSON(http.StatusUnauthorized , gin.H{"error" : "user not logged in"})
		return
	}
	comment := &model.Comment{
		CommentID: commentActionRequest.CommentID,
		UserId:	userID.(int64),
		VideoId: commentActionRequest.VideoID,
		ActionType: commentActionRequest.ActionType,
		CommentText: commentActionRequest.CommentText,
		CreateDate: time.Now().Unix(),
	}
	if commentActionRequest.ActionType == configs.AddComment {
		var sf *snowflake.Node
		if tmp , err := snowflake.NewNode(configs.CommentSnowFlakeNode);err != nil {
			c.JSON(http.StatusInternalServerError , gin.H{"error" : "commentID generate error"})
			return
		}else {
			sf = tmp
		}
		comment.CommentID = sf.Generate().Int64()
		if err := dao.CreateComment(comment); err != nil {
			c.JSON(http.StatusInternalServerError , gin.H{"error" : "create comment error"})
			return
		}
		commentActionResponse.BaseResp = &common.Response{
			StatusCode: http.StatusOK,
			StatusMsg: "create comment successful",
		}
		commentActionResponse.Comment = &common.Comment{
			Id: comment.CommentID,
			User: &common.User{Id:userID.(int64),},
			Content: comment.CommentText,
			CreateDate: strconv.FormatInt(comment.CreateDate, 10),
		}
		c.JSON(http.StatusOK , commentActionResponse)
	}else if commentActionRequest.ActionType == configs.DeleteComment {
		if err := dao.DeleteComment(comment.CommentID); err != nil {
			c.JSON(http.StatusInternalServerError , gin.H{"error" : "delete comment error"})
			return
		}
		c.JSON(http.StatusOK , gin.H{"message" : "delete comment successful"})
	}else {
		commentActionResponse.BaseResp = &common.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg: "invalid comment ActionType",
		}
		c.JSON(http.StatusNotFound , commentActionResponse)
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	var commentListRequest model.CommentListRequest
	var CommentListResponse model.CommentListResponse
	if err := c.ShouldBind(&commentListRequest);err != nil {
		c.JSON(http.StatusNotFound , gin.H{"error" : "comment list req bind error"})
		return
	}
	commentList , err := dao.GetComment(commentListRequest.VideoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError , gin.H{"error" : "Get comment error"})
		return
	}
	for _, v := range commentList {
		timestamp := v.CreateDate
		seconds := timestamp / int64(time.Second)
		nanoseconds := timestamp % int64(time.Second)
		timeObj := time.Unix(seconds, nanoseconds)
		user, err := UserServer.GetUser(v.UserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError , gin.H{"error" : "Get Comment User error"})
		}
		timeStr := timeObj.Format("2006-01-02 15:04:05")
		CommentListResponse.CommentList = append(CommentListResponse.CommentList, &common.Comment{
			Id:         v.CommentID,
			User:       user,
			Content:    v.CommentText,
			CreateDate: timeStr,
		})
	}
	CommentListResponse.BaseResp = &common.Response{
		StatusCode: 0,
		StatusMsg:  "interaction get comment success",
	}
	c.JSON(http.StatusOK , CommentListResponse)
}
