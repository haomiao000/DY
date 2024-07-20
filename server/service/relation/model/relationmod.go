package model
import "main/server/common"
// <------------------------------- Request ------------------------------->
type RelationActionRequest struct {
	UserId     int64 `json:"user_id" form:"user_id" binding:"required"`
	ToUserId   int64 `json:"to_user_id" form:"to_user_id" binding:"required"`
	ActionType int8  `json:"action_type" form:"action_type" binding:"required"`
}
type RelationIdListRequest struct {
	ViewerId int64 `json:"viewer_id" form:"viewer_id"`
	OwnerId  int64 `json:"owner_id" form:"owner_id"`
	Option   int8  `json:"option" form:"option"`
}
type RelationFollowListRequest struct {
	// User id
	UserID int64 `json:"user_id" form:"user_id"`
	// User authentication token
	Token string `json:"token" form:"token"`
}
type RelationFollowerListRequest struct {
	// User id
	UserID int64 `json:"user_id" form:"user_id"`
	// User authentication token
	Token string `json:"token" form:"token"`
}
type RelationFriendListRequest struct {
	// User id
	UserID int64 `json:"user_id" form:"user_id"`
	// User authentication token
	Token string `json:"token" form:"token"`
}
//<------------------------------- gorm -------------------------------> 
type ConcernsInfo struct {
	CommcernsID int64 `gorm:"primarykey;autoIncrement"`
	UserID      int64 `gorm:"column:user_id;	type:INT"`
	FollowerID  int64 `gorm:"column:follower_id;type:INT"`
}
// <------------------------------- Response ------------------------------->
type UserListResponse struct {
	common.Response
	UserList []*common.User `json:"user_list"`
}

func (ConcernsInfo) TableName() string {
	return "concerns_info" 
}
