package oaHandle

import (
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2/officialaccount"
)

// GetAccessToken 获取ak
func GetAccessToken(c *gin.Context, oa *officialaccount.OfficialAccount) {
	token, err := oa.GetAccessToken()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"access_token": token})
}

// GetCallbackIP 获取微信callback IP地址
func GetCallbackIP(c *gin.Context, oa *officialaccount.OfficialAccount) {
	ips, err := oa.GetBasic().GetCallbackIP()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"ip_list": ips})
}

// GetAPIDomainIP 获取微信callback IP地址
func GetAPIDomainIP(c *gin.Context, oa *officialaccount.OfficialAccount) {
	ips, err := oa.GetBasic().GetAPIDomainIP()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"ip_list": ips})
}

// GetAPIDomainIP  清理接口调用次数
func ClearQuota(c *gin.Context, oa *officialaccount.OfficialAccount) {
	err := oa.GetBasic().ClearQuota()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "clear quota success"})
}
