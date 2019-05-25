package main

import (
	"fmt"
	"gin/gin-blog/models"
	"gin/gin-blog/pkg/gredis"
	"gin/gin-blog/pkg/logging"
	"gin/gin-blog/pkg/setting"
	"gin/gin-blog/routers"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://127.0.0.1:8000
// @license.name MIT
// @license.url https://choufengleilei.com
func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
