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


// var UsersLoginInfo = map[string]common.User{
// 	"zhangleidouyin": {
// 		Id:            1,
// 		Name:          "zhanglei",
// 		FollowCount:   10,
// 		FollowerCount: 5,
// 		// IsFollow:      true,
// 	},
// }

func Register(c *gin.Context) {
	var userRegisterReq model.UserRegisterRequest
	if err := c.ShouldBind(&userRegisterReq); err != nil {
		c.JSON(http.StatusNotFound , gin.H{"error" : "missing field"})
		return
	}
	if userRegisterReq.Password == "" || userRegisterReq.Username == "" {
		c.JSON(http.StatusNotFound , gin.H{"error":"Username or Password is empty"})
		return
	}
	// fmt.Println(userRegisterReq.Username)
	// fmt.Println(userRegisterReq.Password)
	var userRegisterInfo model.UserLoginInfo
	if err := dao.FindByUsername(userRegisterReq.Username);err == nil{
		c.JSON(http.StatusConflict, gin.H{"error": "username already exist"})
        return
	}

	userRegisterInfo.Username = userRegisterReq.Username
	userRegisterInfo.Password = middleware.Gen_sha256(userRegisterReq.Password)
	if err := dao.CreateUserLoginInfo(&userRegisterInfo); err!=nil{
		c.JSON(http.StatusInternalServerError , gin.H{"error": "create userLoginInfo error"})
	}
	user := model.User{
		Name:		   userRegisterReq.Username,
		FollowCount:   0,
		FollowerCount: 0,
		// IsFollow:      false,
	}
	if err := dao.CreateUser(&user); err!=nil{
		c.JSON(http.StatusInternalServerError , gin.H{"error": "create user error"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	// fmt.Println("Content-Type:", c.ContentType())
	var userLoginReq model.UserLoginRequest
	if err := c.ShouldBind(&userLoginReq); err != nil {
		c.JSON(http.StatusNotFound , gin.H{"error" : "missing field"})
		return
	}
	if userLoginReq.Username == "" || userLoginReq.Password == "" {
		c.JSON(http.StatusNotFound, gin.H{"error" : "empty login username or password"})
		return
	}

	var userLoginInfo *model.UserLoginInfo
	var err error
	if userLoginInfo , err = dao.CheckUserLoginInfo(&userLoginReq) ; err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
	}
	userID := userLoginInfo.UserID
	token , err := middleware.GenToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "generate token error"})
		return
	}
	// if user , err := GetUser(userID);err != nil{
	// 	c.JSON(http.StatusNotFound , gin.H{"error":"login , get user error"})
	// 	return
	// }else{
	// 	//here use UsersLoginInfo
	// 	UsersLoginInfo[token] = *user
	// }
	var userLoginResp model.UserLoginResponse
	userLoginResp.UserId = userID
	userLoginResp.Token = token
	userLoginResp.StatusCode = 0
	c.JSON(http.StatusOK, userLoginResp)
}
func GetUser(userID int64) (*common.User , error){
	var user *model.User
	var err error
	if user , err = dao.GetUserByUid(userID); err != nil{
		return nil, fmt.Errorf("bind user error maybe userID is wrong: %v", err)
	}
	// fmt.Printf("User: %+v\n", user)
	commonUser := &common.User{
		Id:            user.UserID,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		// IsFollow:      user.IsFollow,
	}
	return commonUser , nil
}
func UserInfo(c *gin.Context) {
	if userID, exists := c.Get("userID"); exists { 
		user , err := GetUser(userID.(int64))
		if err != nil{
			fmt.Println("Failed to get user:", err)
			c.JSON(http.StatusNotFound, model.UserResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "dao can't find user"},
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
