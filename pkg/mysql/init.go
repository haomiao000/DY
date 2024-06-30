package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitMySQL 数据库初始化
func InitMySQL(user, password, ip, port, dbname string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s? charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, dbname)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db.Set("gorm:table_options", "ENDGIN=InnoDB")
	fmt.Println("database init success")
	return nil
}
