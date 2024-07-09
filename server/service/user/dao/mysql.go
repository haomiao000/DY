package dao

import (
	_"fmt"
	"main/internal/initialize"
	"main/middleware"
	"main/server/service/user/model"
)

func FindByUsername(username string) (error){
	var userRegisterInfo model.UserLoginInfo
	err := initialize.DB.Where("username = ?" , username).First(&userRegisterInfo).Error; 
	return err;
}
func CreateUserLoginInfo(userLoginInfo *model.UserLoginInfo) (error){
	err := initialize.DB.Create(&userLoginInfo).Error
	return err;
}
func CreateUser(user *model.User)(error){
	err := initialize.DB.Create(&user).Error
	return err;
}
func CheckUserLoginInfo(userLoginReq *model.UserLoginRequest)(model.UserLoginInfo , error){
	var userLoginInfo model.UserLoginInfo
	err := initialize.DB.Where("username = ? AND password = ?" , userLoginReq.Username , 
	middleware.Gen_sha256(userLoginReq.Password)).First(&userLoginInfo).Error
	return userLoginInfo , err
}
func GetUserByUid(userID int64) (model.User, error){
	var user model.User
	err := initialize.DB.Where("uid = ?" , userID).First(&user).Error
	return user , err
}
