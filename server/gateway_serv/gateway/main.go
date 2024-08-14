package main

import (

	gin "github.com/gin-gonic/gin"
	discovery "github.com/haomiao000/DY/comm/discovery"
	router "github.com/haomiao000/DY/server/gateway_serv/gateway/biz/router"
	initialize "github.com/haomiao000/DY/server/gateway_serv/gateway/initialize"
	rpc "github.com/haomiao000/DY/server/gateway_serv/gateway/initialize/rpc"
)

func main() {
	go initialize.RunMessageServer()
	discovery.Init()
	rpc.Init()
	r := gin.Default()
	router.InitRouter(r)
	r.Run()
}
