package controller

import (
	"fmt"
	"main/internal/initialize"
	"main/middleware"
	"main/pkg/common"
	"main/pkg/model"
	"net/http"
	_ "sync/atomic"

	"github.com/gin-gonic/gin"
)


var usersLoginInfo = map[string]common.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

func Register(c *gin.Context) {
	var user_register_req model.UserRegisterRequest
	user_register_req.Username = c.Query("username")
	user_register_req.Password = c.Query("password")
	// fmt.Println(user_register_req.Username)
	// fmt.Println(user_register_req.Password)
	var user_register_info model.UserLoginInfo
	if err := initialize.DB.Where("username = ?" , user_register_req.Username).First(&user_register_info).Error; 
	err == nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username already exist"})
        return
	}
	user_register_info.Username = user_register_req.Username
	user_register_info.Password = middleware.Gen_sha256(user_register_info.Password)
	if err := initialize.DB.Create(&user_register_info).Error; err!=nil{
		c.JSON(http.StatusInternalServerError , gin.H{"error": "create user_login_info error"})
	}
	user := model.User{
		Name:		   user_register_req.Username,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
	if err := initialize.DB.Create(&user).Error; err!=nil{
		c.JSON(http.StatusInternalServerError , gin.H{"error": "create user error"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var user_login_req model.UserLoginRequest
	user_login_req.Username = c.Query("username")
	user_login_req.Password = c.Query("password")
	var user_login_info model.UserLoginInfo
	if err := initialize.DB.Where("username = ? AND password = ?" , user_login_req.Username , 
	middleware.Gen_sha256(user_login_info.Password)).First(&user_login_info).Error; err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
	}
	userID := user_login_info.UID
	token , err := middleware.GenToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "generate token error"})
		return
	}
	var user_login_resp model.UserLoginResponse
	user_login_resp.UserId = userID
	user_login_resp.Token = token
	user_login_resp.StatusCode = 0
	c.JSON(http.StatusOK, user_login_resp)
}
func GetUser(userID int64) (common.User , error){
	var user model.User
	var err error
	if err = initialize.DB.Where("uid = ?" , userID).First(&user).Error; err != nil{
		fmt.Println("bind user error maybe userID is wrong")
	}
	// fmt.Printf("User: %+v\n", user)
	common_user := common.User{
		Id:            user.UID,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
	return common_user , err
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
				User:     user,
			})
		}
	} else {
		c.JSON(http.StatusInternalServerError, model.UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
