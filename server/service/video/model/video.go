package model

type VideoRecord struct {
	VideoID       int64  `gorm:"column:video_id"`
	UserID        int64  `gorm:"column:user_id"`
	FileName      string `gorm:"column:file_name"`
	UpdateTime    int64  `gorm:"column:update_time"`
	PlayUrl       string `gorm:"column:play_url"`
	CoverUrl      string `gorm:"column:cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
}

func (VideoRecord) TableName() string {
	return "video_records" // TODO 表名
}

type LikeVideo struct {
	UserID  int64 `gorm:"column:user_id"`
	VideoID int64 `gorm:"column:video_id"`
}

func (LikeVideo) TableName() string {
	return "like_videos" // TODO 表名
}
