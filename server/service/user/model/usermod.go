package model

import "main/server/common"

// <------------------------------- Request ------------------------------->
type UserRegisterRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
type UserLoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// <------------------------------- Response ------------------------------->
type UserLoginResponse struct {
	common.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	common.Response
	User common.User `json:"user"`
}

// <------------------------------- gorm ------------------------------->
type UserLoginInfo struct {
	UserID   int64  `gorm:"column:user_id;		 auto_increment;primary_key"`
	Username string `gorm:"column:username; type:nvarchar(255); not null"`
	Password string `gorm:"column:password; type:nvarchar(255); not null"`
}
type User struct {
	UserID        int64  `gorm:"column:user_id;			 auto_increment;primary_key;"`
	Name          string `gorm:"column:name;		 type:varchar(100);not null"`
	FollowCount   int64  `gorm:"column:followcount;  type:INT"`
	FollowerCount int64  `gorm:"column:followercount;type:INT"`
	IsFollow      bool   `gorm:"column:isfollow;	 type:bool"`
}

func (UserLoginInfo) TableName() string {
	return "user_login_info"
}
func (User) TableName() string {
	return "user"
}
