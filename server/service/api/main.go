package main

import (
	rpc "main/server/service/api/initialize/rpc"
	initialize "main/server/service/api/initialize"
	router "main/server/service/api/biz/router"
	gin "github.com/gin-gonic/gin"
)

func main() {
	go initialize.RunMessageServer()
	rpc.Init()
	r := gin.Default()
	router.InitRouter(r)
	r.Run()
}