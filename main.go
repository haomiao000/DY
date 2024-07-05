package main

import (
	"github.com/gin-gonic/gin"
	"main/internal/initialize"
)

func main() {
	go initialize.InitServer()

	r := gin.Default()

	// TODO 数据库初始化
	// if err := db.InitMySQL(user,password,ip,port,dbname); err != nil {
	// 	fmt.Printf("database init error: %v",err)
	// 	return
	// }

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
