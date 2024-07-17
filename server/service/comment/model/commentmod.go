package model
import "main/server/common"
//<------------------------------- Request ------------------------------->
type CommentActionRequest struct {
	Token string `json:"token" form:"token" binding:"required"`
	VideoID int64 `json:"video_id" form:"video_id" binding:"required"`
	ActionType int8 `json:"action_type" form:"action_type" binding:"required"`
	CommentText string `json:"comment_text" form:"comment_text"`
	// The comment id to be deleted is used when action_type=2
	CommentID int64 `json:"comment_id" form:"comment_id"`
}
type CommentListRequest struct {
	Token string `json:"token" form:"token" binding:"required"`
	VideoID int64 `json:"video_id" form:"video_id" binding:"required"`
}
//<------------------------------- Response ------------------------------->
type CommentListResponse struct {
	BaseResp *common.Response 
	CommentList []*common.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	BaseResp *common.Response `json:"base_resp"`
	Comment *common.Comment `json:"comment,omitempty"`
}
//<------------------------------- gorm ------------------------------->

type Comment struct {
	CommentID   int64  `gorm:"column:comment_id;	primary_key;"`
	UserId      int64  `gorm:"column:user_id;		type:INT;not null"`
	VideoId     int64  `gorm:"column:video_id;		type:INT;not null"`
	ActionType  int8   `gorm:"column:action_type;	type:tinyint;not null"`
	CommentText string `gorm:"column:comment_text; 	type:varchar(256);not null"`
	CreateDate  int64  `gorm:"column:create_time;	type:INT;not null"`
}
func (Comment) TableName() string {
	return "comment" 
}
