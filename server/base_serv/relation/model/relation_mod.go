package model

type ConcernsInfo struct {
	ConcernsID int64 `gorm:"primarykey;autoIncrement"`
	UserID      int64 `gorm:"column:user_id;	type:INT"`
	FollowerID  int64 `gorm:"column:follower_id;type:INT"`
}
func (ConcernsInfo) TableName() string {
	return "concerns_info" 
}