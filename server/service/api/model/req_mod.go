package model

//@Register
type UserRegisterRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
//@Login
type UserLoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
//@UserInfo
type UserInfoRequest struct {
	UserID    int64 `json:"user_id" form:"user_id"`
	Token     string  `json:"token" form:"token"`
}
//@Favorite
type FavoriteActionRequest struct {
	Token      string `json:"token" form:"token" binding:"required"`
	VideoID    int64  `json:"video_id" form:"video_id" binding:"required"`
	ActionType int8   `json:"action_type" form:"action_type" binding:"required"`
}
//@FavoriteList
type FavoriteListRequest struct {
	UserID 	   int64 `json:"user_id" form:"user_id" binding:"required"`
	Token      string `json:"token" form:"token" binding:"required"`
}
//@Comment
type CommentActionRequest struct {
	Token       string `json:"token" form:"token" binding:"required"`
	VideoID     int64  `json:"video_id" form:"video_id" binding:"required"`
	ActionType  int8   `json:"action_type" form:"action_type" binding:"required"` // 1 or 2
	CommentText string `json:"comment_text" form:"comment_text"`
	CommentID 	int64  `json:"comment_id" form:"comment_id"`
}
//@CommentList
type CommentListRequest struct {
	Token   string `json:"token" form:"token" binding:"required"`
	VideoID int64  `json:"video_id" form:"video_id" binding:"required"`
}
//@RelationAction
type RelationActionRequest struct {
	Token      string `json:"token" form:"token" binding:"required"`
	ToUserId   int64 `json:"to_user_id" form:"to_user_id" binding:"required"`
	ActionType int8  `json:"action_type" form:"action_type" binding:"required"`
}
//@FollowList
type RelationFollowListRequest struct {
	UserID int64 `json:"user_id" form:"user_id"`
	Token string `json:"token" form:"token"`
}
//@FollowerList
type RelationFollowerListRequest struct {
	UserID int64 `json:"user_id" form:"user_id"`
	Token string `json:"token" form:"token"`
}
//@FriendList
type RelationFriendListRequest struct {
	UserID int64 `json:"user_id" form:"user_id"`
	Token string `json:"token" form:"token"`
}