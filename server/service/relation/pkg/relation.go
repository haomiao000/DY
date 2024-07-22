package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"main/server/common"
	// "main/test/testcase"
	"main/server/service/relation/model"
	"main/server/service/relation/dao"
	"main/configs"
)


// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	var relationActionRequest model.RelationActionRequest
	if err := c.ShouldBind(&relationActionRequest); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "binding relationAction Request error"})
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not logged in"})
		return
	}
	if relationActionRequest.ActionType != configs.Follow && relationActionRequest.ActionType != configs.UnFollow {
		c.JSON(http.StatusInternalServerError , common.Response{StatusCode: 1 , StatusMsg: "invalid action type"})
		return
	}
	if userID.(int64) == relationActionRequest.ToUserId {
		c.JSON(http.StatusInternalServerError , common.Response{StatusCode: 1 , StatusMsg: "user can not followe itself"})
		return
	}
	followStatus , err := dao.FindRelationInfo(userID.(int64) , relationActionRequest.ToUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError , common.Response{StatusCode: 1 , StatusMsg: "relation action occurs ErrRecordNotFound"})
		return
	}
	if followStatus == configs.IsFollow {
		if relationActionRequest.ActionType == configs.Follow {
			c.JSON(http.StatusOK , common.Response{StatusCode: 0 , StatusMsg: "u have follow the author"})
			return
		} else {
			if err := dao.DeleteRelationInfo(userID.(int64) , relationActionRequest.ToUserId); err != nil {
				c.JSON(http.StatusInternalServerError , common.Response{StatusCode: 1 , StatusMsg: "delete relation info error"})
				return
			}
			c.JSON(http.StatusOK , common.Response{StatusCode: 1 , StatusMsg: "unfollow successful"})
		}
	}else {
		if relationActionRequest.ActionType == configs.UnFollow {
			c.JSON(http.StatusOK , common.Response{StatusCode: 0 , StatusMsg: "u have not follow the author"})
			return
		}else {
			if err := dao.CreateRelationInfo(userID.(int64) , relationActionRequest.ToUserId); err != nil {
				c.JSON(http.StatusInternalServerError , common.Response{StatusCode: 1 , StatusMsg: "create relation info error"})
				return
			}
			c.JSON(http.StatusOK , common.Response{StatusCode: 0 , StatusMsg: "follow the author successful"})
		}
	}
}

//follows
func FollowList(c *gin.Context) {
	var relationFollowListRequest model.RelationFollowListRequest
	if err := c.ShouldBind(&relationFollowListRequest); err != nil {
		c.JSON(http.StatusNotFound , model.UserListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg: "bind follow user list req error",
			},	
			UserList: nil,
		})
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not logged in"})
		return
	}
	mp := *GetFollowMap(userID.(int64) , c)
	var followList []*common.User
	if userList , err := dao.GetFollowUserList(relationFollowListRequest.UserID); err != nil {
		c.JSON(http.StatusInternalServerError , model.UserListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg: "get follow user list error",
			},	
			UserList: nil,
		})
		return
	}else {
		for _ , o := range userList {
			followList = append(followList, &common.User{
				Id: o.UserID,
				Name: o.Name,
				FollowCount: o.FollowCount,
				FollowerCount: o.FollowerCount,
				IsFollow: mp[o.UserID],
			})
		}
	}
	c.JSON(http.StatusOK , model.UserListResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg: "get follow successful",
		},
		UserList: followList,
	})
}
//funs
func FollowerList(c *gin.Context) {
	var relationFollowerListRequest model.RelationFollowerListRequest
	if err := c.ShouldBind(&relationFollowerListRequest); err != nil {
		c.JSON(http.StatusNotFound , model.UserListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg: "bind follower user list req error",
			},
			UserList: nil,
		})
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not logged in"})
		return
	}
	mp := *GetFollowMap(userID.(int64) , c)
	var followerList []*common.User 
	if userList , err := dao.GetFollowerUserList(relationFollowerListRequest.UserID); err != nil {
		c.JSON(http.StatusInternalServerError , model.UserListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg: "get follower user list error",
			},	
			UserList: nil,
		})
		return
	}else {
		for _ , o := range userList {
			followerList = append(followerList, &common.User{
				Id: o.UserID,
				Name: o.Name,
				FollowCount: o.FollowCount,
				FollowerCount: o.FollowerCount,
				IsFollow: mp[o.UserID],
			})
		}
	}
	c.JSON(http.StatusOK , model.UserListResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg: "get follow successful",
		},
		UserList: followerList,
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	var relationFriendListRequest model.RelationFriendListRequest
	if err := c.ShouldBind(&relationFriendListRequest); err != nil {
		c.JSON(http.StatusOK , model.UserListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg: "bind friend list req error",
			},
			UserList: nil,
		})
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not logged in"})
		return
	}
	mp := *GetFollowMap(userID.(int64) , c)
	var respUser []*common.User
	if friendList , err := dao.GetMutualFollowers(relationFriendListRequest.UserID); err != nil {
		c.JSON(http.StatusOK , model.UserListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg: "get friend list error",
			},
		})
		return
	}else {
		for _ , o := range friendList {
			respUser = append(respUser, &common.User{
				Id: o.UserID,
				Name: o.Name,
				FollowCount: o.FollowCount,
				FollowerCount: o.FollowerCount,
				IsFollow: mp[o.UserID],
			})
		}
	}
	c.JSON(http.StatusOK , model.UserListResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg: "get friend list successful",
		},
		UserList: respUser,
	})
}
