package main

import (
	"flag"
	"fmt"
	"log"
	"wxapi-go/config"
	"wxapi-go/message"
	"wxapi-go/util"

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
		util.NewWxConfig(&conf.Wx)
	} else {
		log.Println("Error:config file is nil")
	}

	gin.SetMode(ginmode)
	router := gin.Default()

	router.GET("/", util.WXCheckSignature)
	router.POST("/", message.WXMsgReceive)

	log.Fatalln(router.Run(":" + port))

	// // 调用 GetAccessToken 获取 Access Token
	// accessToken, ExpiresIn, err := util.GetAccessToken(appID, appSecret)
	// if err != nil {
	// 	log.Println("Error getting Access Token:", err)
	// }
	// util.Access_Token = accessToken //更新accessToken
	// fmt.Println("Now the Access Token:", accessToken, " ExpiresIn:", ExpiresIn)

	// // 启动一个定时任务，每隔 2 小时执行一次获取 Access Token 的操作
	// ticker := time.NewTicker(2 * time.Hour)
	// go func() {
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			// 调用 GetAccessToken 获取 Access Token
	// 			accessToken, ExpiresIn, err := util.GetAccessToken(appID, appSecret)
	// 			if err != nil {
	// 				log.Println("Error getting Access Token:", err)
	// 				continue
	// 			}
	// 			util.Access_Token = accessToken //更新accessToken
	// 			fmt.Println("Now the Access Token:", accessToken, " ExpiresIn:", ExpiresIn)
	// 		}
	// 	}
	// }()
	// // 阻塞主程序，保持定时任务持续运行
	// select {}
}
