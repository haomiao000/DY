package initialize

import (
	gorm "gorm.io/gorm"
	logger "gorm.io/gorm/logger"
	configs "github.com/haomiao000/DY/server/service/user/configs"
	logrus "gorm.io/plugin/opentelemetry/logging/logrus"
	mysql "gorm.io/driver/mysql"
	schema "gorm.io/gorm/schema"
	"fmt"
	"time"
)
func InitDB() *gorm.DB {
	dsn := fmt.Sprintf(configs.MySqlDSN, configs.UserDBUser, configs.UserDBPassword, configs.UserDBIP, configs.UserDBPort, configs.UserDBName)
	newLogger := logger.New(
		logrus.NewWriter(), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL Threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // Disable color printing
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		fmt.Printf("[DB Err]\t%v\n", err)
	}
	return db
}