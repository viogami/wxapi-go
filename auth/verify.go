package auth

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
	"sort"
	"strings"
	"wxapi-go/config"

	"github.com/gin-gonic/gin"
)

var (
	wxToken   string
	appID     string
	appSecret string
)

// 微信公众号签名检查
func CheckSignature(signature, timestamp, nonce, token string) bool {
	arr := []string{timestamp, nonce, token}
	// 字典序排序
	sort.Strings(arr)

	n := len(timestamp) + len(nonce) + len(token)
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < len(arr); i++ {
		b.WriteString(arr[i])
	}

	return Sha1(b.String()) == signature
}

// 进行Sha1编码
func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 微信接入校验
func WXCheckSignature(c *gin.Context) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	ok := CheckSignature(signature, timestamp, nonce, wxToken)
	if !ok {
		log.Println("微信公众号签名检查失败!")
		return
	}

	log.Println("微信公众号接入校验成功!")
	_, _ = c.Writer.WriteString(echostr)
}

func NewWxConfig(conf *config.Config) {
	wxToken = conf.OfficialAccount.Token
	appID = conf.OfficialAccount.AppID
	appSecret = conf.OfficialAccount.AppSecret
}
