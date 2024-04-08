package server

import (
	"context"
	"fmt"
	"wxapi-go/auth"
	"wxapi-go/config"

	"github.com/gin-gonic/gin"
	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
)

var oa *officialaccount.OfficialAccount

type Server struct {
	engine *gin.Engine
	wx     *wechat.Wechat
	conf   *config.Config
	redis  *cache.Redis
}

func New(conf *config.Config) (*Server, error) {
	gin.SetMode(conf.Server.Env)
	r := gin.Default()

	redisCache := cache.NewRedis(context.Background(), &conf.Redis)

	wx := wechat.NewWechat()

	return &Server{
		engine: r,
		wx:     wx,
		redis:  redisCache,
		conf:   conf,
	}, nil
}

func (s *Server) RunOA() error {
	s.SetupRoute()
	// 设置redis缓存
	s.wx.SetCache(s.redis)
	// 初始化微信公众号实例
	oa = s.wx.GetOfficialAccount(&s.conf.OfficialAccount)
	// 初始化微信认证配置
	auth.NewWxConfig(s.conf)
	//读取服务器地址
	addr := fmt.Sprintf("%s:%d", s.conf.Server.Address, s.conf.Server.Port)

	return s.engine.Run(addr)
}
