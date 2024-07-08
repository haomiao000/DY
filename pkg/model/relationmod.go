package model
import "main/pkg/common"

type UserListResponse struct {
	common.Response
	UserList []common.User `json:"user_list"`
}
//<------------------------------- gorm ------------------------------->