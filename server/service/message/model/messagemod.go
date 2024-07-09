package model
import "main/server/common"
type ChatResponse struct {
	common.Response
	MessageList []common.Message `json:"message_list"`
}

//<------------------------------- gorm ------------------------------->