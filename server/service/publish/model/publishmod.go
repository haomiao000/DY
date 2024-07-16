package model

import "main/server/common"

type VideoListResponse struct {
	BaseResp  *common.Response
	VideoList []*common.Video `json:"video_list"`
}
//<------------------------------- gorm ------------------------------->