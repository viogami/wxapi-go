package server

import (
	"wxapi-go/auth"
	"wxapi-go/oaHandle"

	"github.com/gin-gonic/gin"
)

func (s *Server) SetupRoute() *gin.Engine {
	r := s.engine

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello, Welcome to wxapi-go ,this is the backend of the '未来bot研究所'")
	})
	r.GET("/v1", func(c *gin.Context) { c.Redirect(302, "/") })

	r.GET("/v1/oa", auth.WXCheckSignature)
	r.POST("/v1/oa", oaHandle.WXMsgReceive)
	r.GET("/v1/oa/menu", func(ctx *gin.Context) { oaHandle.CheckMenu(ctx, oa) })

	//获取ak
	r.GET("/v1/oa/basic/get_access_token", func(ctx *gin.Context) { oaHandle.GetAccessToken(ctx, oa) })
	//获取微信callback IP
	r.GET("/v1/oa/basic/get_callback_ip", func(ctx *gin.Context) { oaHandle.GetCallbackIP(ctx, oa) })
	//获取微信API接口 IP
	r.GET("/v1/oa/basic/get_api_domain_ip", func(ctx *gin.Context) { oaHandle.GetAPIDomainIP(ctx, oa) })
	//清理接口调用次数
	r.GET("/v1/oa/basic/clear_quota", func(ctx *gin.Context) { oaHandle.ClearQuota(ctx, oa) })

	return r
}
