package message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Message struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

func Postmsg(msgtext string, accessToken *string) {
	// 实例化 Message 结构体
	message := Message{
		ToUser:  "OPENID",
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: msgtext,
		},
	}
	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Fatal("Error encoding JSON data:", err)
	}

	// 构建请求
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", *accessToken)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Error creating HTTP request:", err)
	}

	defer resp.Body.Close()

	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error to read the response body", err)
		return
	}
	//输出响应内容
	log.Println(string(respbody))

}
