package model

type VideoRecord struct {
	VideoID       int64  `gorm:"column:video_id; type:int; auto_increment;primary_key"`
	UserID        int64  `gorm:"column:user_id; type:int"`
	FileName      string `gorm:"column:file_name; type:nvarchar(255); not null"`
	UpdateTime    int64  `gorm:"column:update_time; type:int"`
	PlayUrl       string `gorm:"column:play_url; type:nvarchar(255);"`
	CoverUrl      string `gorm:"column:cover_url; type:nvarchar(255); not null"`
	FavoriteCount int64  `gorm:"column:favorite_count; type:int"`
	CommentCount  int64  `gorm:"column:comment_count; type:int"`
}

func (VideoRecord) TableName() string {
	return "video_records" // TODO 表名
}

type Favorite struct {
	UserID     int64 `gorm:"column:user_id; 	type:INT"`
	VideoID    int64 `gorm:"column:video_id;	type:INT"`
	CreateDate int64 `gorm:"column:create_time;	type:INT;not null"`
}
