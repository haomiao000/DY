package initialize

import (
	"fmt"
	"time"

	configs "github.com/haomiao000/DY/server/base_serv/video/configs"
	mysql "gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
	logger "gorm.io/gorm/logger"
	schema "gorm.io/gorm/schema"
	logrus "gorm.io/plugin/opentelemetry/logging/logrus"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf(configs.MySqlDSN, configs.VideorDBUser, configs.VideoDBPassword, configs.VideoDBIP, configs.VideoDBPort, configs.VideoDBName)
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
