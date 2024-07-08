package model
import "main/server/common"

//<------------------------------- Request ------------------------------->
type UserRegisterRequest struct{
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserLoginRequest struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

//<------------------------------- Response ------------------------------->
type UserLoginResponse struct {
	common.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	common.Response
	User common.User `json:"user"`
}

//<------------------------------- gorm ------------------------------->
type UserLoginInfo struct {
	UID      	  int64  `gorm:"column:uid;		 auto_increment;primary_key"`
	Username      string `gorm:"column:username; type:nvarchar(255); not null"`
	Password 	  string `gorm:"column:password; type:nvarchar(255); not null"`
}
type User struct {
	UID           int64  `gorm:"column:uid;			 auto_increment;primary_key;"`
	Name          string `gorm:"column:name;		 type:varchar(100);not null"`
	FollowCount   int64  `gorm:"column:followcount;  type:INT"`
	FollowerCount int64  `gorm:"column:followercount;type:INT"`
	IsFollow      bool   `gorm:"column:isfollow;	 type:bool"`
}
func (UserLoginInfo) TableName() string{
	return "user_login_info"
}
func (User) TableName() string{
	return "user"
}