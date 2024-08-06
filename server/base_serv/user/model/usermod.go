package model

type UserLoginInfo struct {
	UserID   int64  `gorm:"column:user_id;		 auto_increment;primary_key"`
	Username string `gorm:"column:username; type:nvarchar(255); not null"`
	Password string `gorm:"column:password; type:nvarchar(255); not null"`
}
type User struct {
	UserID        int64  `gorm:"column:user_id;		 auto_increment;primary_key;"`
	Name          string `gorm:"column:name;		 type:varchar(100);not null"`
	FollowCount   int64  `gorm:"column:followcount;  type:INT"`
	FollowerCount int64  `gorm:"column:followercount;type:INT"`
}

func (UserLoginInfo) TableName() string {
	return "user_login_info"
}
func (User) TableName() string {
	return "user"
}
