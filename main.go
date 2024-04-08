package main

import (
	"flag"
	"log"
	"wxapi-go/config"
	"wxapi-go/server"
)

func main() {
	appConfig := flag.String("config", "config/app.yaml", "application config path")
	conf, err := config.ConfigParse(appConfig)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	s, err := server.New(conf)
	if err != nil {
		log.Fatalf("Init server failed: %v", err)
	}

	if err := s.RunOA(); err != nil {
		log.Fatalf("Run server failed: %v", err)
	}
	log.Println("Server start success... port:", conf.Server.Port, " mode:", conf.Server.Env)

	// 调用 StartAccessTokenScheduler 开始定时任务
	//go auth.StartAccessTokenScheduler(conf.OfficialAccount.AppID, conf.OfficialAccount.AppSecret)
}
