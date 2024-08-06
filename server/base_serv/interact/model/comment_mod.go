package model

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
