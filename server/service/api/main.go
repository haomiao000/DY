package main

import (
	rpc "github.com/haomiao000/DY/server/service/api/initialize/rpc"
	initialize "github.com/haomiao000/DY/server/service/api/initialize"
	router "github.com/haomiao000/DY/server/service/api/biz/router"
	gin "github.com/gin-gonic/gin"
)

func main() {
	go initialize.RunMessageServer()
	rpc.Init()
	r := gin.Default()
	router.InitRouter(r)
	r.Run()
}