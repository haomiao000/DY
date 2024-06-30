package model

import "main/pkg/common"

type VideoListResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list"`
}