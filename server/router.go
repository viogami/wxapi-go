package server

import (
	"wxapi-go/auth"
	"wxapi-go/controller/message"

	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	router := gin.Default()

	router.GET("/", auth.WXCheckSignature)
	router.POST("/", message.WXMsgReceive)

	return router
}
