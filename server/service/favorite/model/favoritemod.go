package model
//<------------------------------- Request ------------------------------->
type FavoriteActionRequest struct {
    Token      string `json:"token" form:"token" binding:"required"`
    VideoID    int64  `json:"video_id" form:"video_id" binding:"required"`
    ActionType int8   `json:"action_type" form:"action_type" binding:"required"`
}
// type FavoriteListRequest struct {
// 	// User id
// 	UserID int64 `json:"user_id" form:"user_id"`
// 	// User authentication token
// 	Token string `json:"token" form:"token"`
// }
//<------------------------------- gorm ------------------------------->
type Favorite struct {
	UserID     int64 `gorm:"column:user_id; 	type:INT"`
	VideoID    int64 `gorm:"column:video_id;	type:INT"`
	ActionType int8  `gorm:"column:action_type;	type:tinyint;not null"`
	CreateDate int64 `gorm:"column:create_time;	type:INT;not null"`
}
