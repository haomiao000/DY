package model

type Favorite struct {
	UserID     int64 `gorm:"column:user_id; 	type:INT"`
	VideoID    int64 `gorm:"column:video_id;	type:INT"`
	CreateDate int64 `gorm:"column:create_time;	type:INT;not null"`
}
func (Favorite) TableName() string {
	return "favorite" 
}

