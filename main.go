package main

import (
	"flag"
	"fmt"
	"log"
	"wxapi-go/auth"
	"wxapi-go/config"
	"wxapi-go/server"

	"github.com/gin-gonic/gin"
)

var (
	port    string
	ginmode string
)

func main() {
	appConfig := flag.String("config", "config/app.yaml", "application config path")
	conf, _ := config.ConfigParse(appConfig)
	if conf != nil {
		port = fmt.Sprint(conf.Server.Port)
		ginmode = conf.Server.Env
		auth.NewWxConfig(&conf.Wx)
	} else {
		log.Fatalln("Error:config file is nil")
	}

	gin.SetMode(ginmode)

	r := server.SetupRoute()
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal("Server start failed...")
	}
	log.Println("Server start success... port:", port, " mode:", ginmode)

	// 调用 StartAccessTokenScheduler 开始定时任务
	auth.StartAccessTokenScheduler(conf.Wx.AppID, conf.Wx.AppSecret)
}
