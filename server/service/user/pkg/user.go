package pkg

import (
	"fmt"
	"main/middleware"
	"main/server/common"
	"main/server/service/user/model"
	"main/server/service/user/dao"
	"net/http"
	_ "sync/atomic"

	"github.com/gin-gonic/gin"
)


var UsersLoginInfo = map[string]common.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

func Register(c *gin.Context) {
	var userRegisterReq model.UserRegisterRequest
	userRegisterReq.Username = c.Query("username")
	userRegisterReq.Password = c.Query("password")
	// fmt.Println(userRegisterReq.Username)
	// fmt.Println(userRegisterReq.Password)
	var userRegisterInfo model.UserLoginInfo
	if err := dao.FindByUsername(userRegisterReq.Username);err == nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username already exist"})
        return
	}
	userRegisterInfo.Username = userRegisterReq.Username
	userRegisterInfo.Password = middleware.Gen_sha256(userRegisterInfo.Password)
	if err := dao.CreateUserLoginInfo(&userRegisterInfo); err!=nil{
		c.JSON(http.StatusInternalServerError , gin.H{"error": "create userLoginInfo error"})
	}
	user := model.User{
		Name:		   userRegisterReq.Username,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
	if err := dao.CreateUser(&user); err!=nil{
		c.JSON(http.StatusInternalServerError , gin.H{"error": "create user error"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var userLoginReq model.UserLoginRequest
	userLoginReq.Username = c.Query("username")
	userLoginReq.Password = c.Query("password")
	var userLoginInfo model.UserLoginInfo
	var err error
	if userLoginInfo , err = dao.CheckUserLoginInfo(&userLoginReq) ; err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
	}
	userID := userLoginInfo.UID
	token , err := middleware.GenToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "generate token error"})
		return
	}
	var userLoginResp model.UserLoginResponse
	userLoginResp.UserId = userID
	userLoginResp.Token = token
	userLoginResp.StatusCode = 0
	c.JSON(http.StatusOK, userLoginResp)
}
func GetUser(userID int64) (*common.User , error){
	var user model.User
	var err error
	if user , err = dao.GetUserByUid(userID); err != nil{
		fmt.Println("bind user error maybe userID is wrong")
	}
	// fmt.Printf("User: %+v\n", user)
	commonUser := &common.User{
		Id:            user.UID,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
	return commonUser , err
}
func UserInfo(c *gin.Context) {
	if uid, exists := c.Get("uid"); exists { 
		user , err := GetUser(uid.(int64))
		if err != nil{
			fmt.Println("Failed to get user:", err)
			c.JSON(http.StatusInternalServerError, model.UserResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "User can't find"},
			})
		}else {
			c.JSON(http.StatusOK, model.UserResponse{
				Response: common.Response{StatusCode: 0},
				User:     *user,
			})
		}
	} else {
		c.JSON(http.StatusInternalServerError, model.UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
