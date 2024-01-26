package message

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
	"wxapi-go/util"

	"github.com/gin-gonic/gin"
)

// WXTextMsg 微信文本消息结构体
type WXTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	MsgId        int64
}

// WXMsgReceive 微信消息接收
func WXMsgReceive(c *gin.Context) {
	var textMsg WXTextMsg
	err := c.ShouldBindXML(&textMsg)
	if err != nil {
		log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
		return
	}
	//log.Printf("[消息接收] - 收到消息, 消息类型为: %s, 消息内容为: %s\n", textMsg.MsgType, textMsg.Content)

	switch textMsg.MsgType {
	case "event":
		simpleReply(c, textMsg.ToUserName, textMsg.FromUserName, "感谢你的关注,我是聊天bot:vio🥰使用Gpt3.5turbo的接口，发送消息即可对话。") // 关注后的默认回复

	case "text":
		GptReplyWXMsg(c, textMsg.ToUserName, textMsg.FromUserName, textMsg.Content, &util.Access_Token) // 调用gpt回复

	default:
		simpleReply(c, textMsg.ToUserName, textMsg.FromUserName, "残念！这个消息类型我的开发者还没有进行相应的设置😭") // 关注后的默认回复
	}
}

// WXRepTextMsg 微信回复文本消息结构体
type WXRepTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	// 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName xml.Name `xml:"xml"`
}

// GptReplyWXMsg 微信消息回复
func GptReplyWXMsg(c *gin.Context, fromUser, toUser, usercontent string, accessToken *string) {
	//定义post的url地址
	URL := "https://bot.masterkagami.com/post"
	// 准备POST请求的数据
	formData := url.Values{
		"usermsg": {usercontent},
	}
	// 发送POST请求
	resp, err := http.PostForm(URL, formData)
	if err != nil {
		log.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error to read the response body", err)
		return
	}
	//输出响应内容
	log.Println(string(respbody))

	repTextMsg := WXRepTextMsg{
		ToUserName:   toUser,
		FromUserName: fromUser,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      string(respbody),
	}

	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[ERROR] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(msg)
}

func simpleReply(c *gin.Context, fromUser, toUser, text string) {
	repTextMsg := WXRepTextMsg{
		ToUserName:   toUser,
		FromUserName: fromUser,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      text,
	}

	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("[ERROR] - 将对象进行XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(msg)
}
