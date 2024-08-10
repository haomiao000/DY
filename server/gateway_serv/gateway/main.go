package main

import (
	gin "github.com/gin-gonic/gin"
	router "github.com/haomiao000/DY/server/gateway_serv/gateway/biz/router"
	initialize "github.com/haomiao000/DY/server/gateway_serv/gateway/initialize"
	rpc "github.com/haomiao000/DY/server/gateway_serv/gateway/initialize/rpc"
	// redis "github.com/haomiao000/DY/comm/redis"
)

func main() {
	go initialize.RunMessageServer()
	// redis.Init()
	rpc.Init()
	r := gin.Default()
	router.InitRouter(r)
	r.Run()
}
