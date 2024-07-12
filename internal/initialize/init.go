package initialize

import (
	"fmt"
	"log"
	"os"

	"main/configs"
	"main/internal"
	"main/server/service/user/model"
	videomodel "main/server/service/video/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitLog() error {
	logFile, err := os.OpenFile("database.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("无法打开日志文件: %v\n", err)
		return err
	}
	defer logFile.Close()
	return nil
}

func InitMySQL() error {
	logFile, _ := internal.SetLogFile()
	// 设置日志输出到文件
	log.SetOutput(logFile)
	var err error
	// 构建 DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", configs.DBUser, configs.DBPassword, configs.DBIP, configs.DBPort, configs.DBName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("[DB Err]\t%v\n", err)
		return err
	}
	if err = DB.AutoMigrate(&model.UserLoginInfo{}); err != nil {
		fmt.Printf("[DB Err]\t%v\n", err)
		return err
	}
	if err = DB.AutoMigrate(&model.User{}); err != nil {
		fmt.Printf("[DB Err]\t%v\n", err)
		return err
	}
	if err = DB.AutoMigrate(&videomodel.VideoRecord{}); err != nil {
		fmt.Printf("[DB Err]\t%v\n", err)
		return err
	}
	if err = DB.AutoMigrate(&videomodel.LikeVideo{}); err != nil {
		fmt.Printf("[DB Err]\t%v\n", err)
		return err
	}
	log.Println("数据库初始化成功")
	return nil
}

func InitServer() error {
	err1 := InitLog()
	if err1 != nil {
		fmt.Println("Init Log error")
		return err1
	}
	err2 := InitMySQL()
	if err2 != nil {
		fmt.Println("Init Mysql error")
		return err2
	}
	// err3 :=
	RunMessageServer()
	return nil
}
