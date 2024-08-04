package dao

import (
	gorm "gorm.io/gorm"
	model "main/server/service/user/model"
	rpc_user "main/server/grpc_gen/rpc_user"
	middleware "main/server/common/middleware"
	configs "main/server/common/configs"
	"errors"
	"fmt"
)
type MysqlManager struct {
	userDB *gorm.DB
	userLoginDB *gorm.DB
}

func (u MysqlManager) CreateUserLoginInfo(userLoginInfo *model.UserLoginInfo) error {
	var temp model.UserLoginInfo
	err := u.userLoginDB.Where("username = ?", userLoginInfo.Username).First(&temp).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		fmt.Printf("mysql select failed,%s\n", err)
		return err
	}
	if temp.Username != "" {
		err = errors.New(configs.MysqlAlreadyExists)
		return err
	}
	err = u.userLoginDB.Create(&userLoginInfo).Error
	if err != nil {
		fmt.Printf("mysql insert failed,%s\n", err)
		return err
	}
	return nil
}

func (u MysqlManager) CreateUser(user *model.User) error {
	err := u.userDB.Create(&user).Error
	if err != nil {
		fmt.Printf("mysql insert failed,%s\n", err)
		return err
	}
	return nil
}

func (u MysqlManager) CheckUserLoginInfo(userLoginReq *rpc_user.UserLoginRequest)(*model.UserLoginInfo , error) {
	var userLoginInfo model.UserLoginInfo
	err := u.userLoginDB.Where("username = ? AND password = ?" , userLoginReq.Username , 
	middleware.Gen_sha256(userLoginReq.Password)).First(&userLoginInfo).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Printf("username or password is error ,%s\n", err)
		return nil , err
	}
	if err != nil {
		fmt.Printf("mysql select failed,%s\n", err)
		return nil , err
	}
	return &userLoginInfo , nil
}

func (u MysqlManager) GetUserByUid(userID int64) (*model.User, error) {
	var user model.User
	err := u.userDB.Where("user_id = ?" , userID).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Printf("user id is not exist in db,%s\n", err)
		return nil , err
	}
	if err != nil {
		fmt.Printf("mysql select failed,%s\n", err)
		return nil , err
	}
	return &user , nil
}

func (u MysqlManager) GetUserListByUserId(userID []int64) ([]*model.User , error) {
	var users []*model.User
	// Query the database for the given user IDs
	err := u.userDB.Where("user_id IN ?", userID).Find(&users).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Printf("user is not exist in db,%s\n", err)
		return nil , err
	}
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewMysqlManager(db *gorm.DB) *MysqlManager {
	m := db.Migrator()
	if !m.HasTable(&model.UserLoginInfo{}) {
		err := m.CreateTable(&model.UserLoginInfo{})
		if err != nil {
			fmt.Printf("create mysql table failed,%s\n", err)
		}
	}
	if !m.HasTable(&model.User{}) {
		err := m.CreateTable(&model.User{})
		if err != nil {
			fmt.Printf("create mysql table failed,%s\n", err)
		}
	}
	return &MysqlManager{
		userDB: db,
		userLoginDB: db,
	}
}
