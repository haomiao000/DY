package model
import "main/server/common"

type UserListResponse struct {
	common.Response
	UserList []common.User `json:"user_list"`
}
//<------------------------------- gorm ------------------------------->