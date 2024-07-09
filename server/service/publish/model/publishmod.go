package model

import "main/server/common"

type VideoListResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list"`
}
//<------------------------------- gorm ------------------------------->