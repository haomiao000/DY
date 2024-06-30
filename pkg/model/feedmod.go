package model
import "main/pkg/common"

type FeedResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}