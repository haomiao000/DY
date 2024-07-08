package model
import "main/pkg/common"
type ChatResponse struct {
	common.Response
	MessageList []common.Message `json:"message_list"`
}

//<------------------------------- gorm ------------------------------->