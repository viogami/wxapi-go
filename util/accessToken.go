package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var Access_Token string //access_token全局变量

// AccessTokenResponse 用于解析获取 Access Token 的响应
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetAccessToken 获取微信公众号的 Access Token
func GetAccessToken(appID, appSecret string) (string, int, error) {
	// 构建请求 URL
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appID, appSecret)

	// 发送 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}

	// 解析 JSON 响应
	var accessTokenResponse AccessTokenResponse
	err = json.Unmarshal(body, &accessTokenResponse)
	if err != nil {
		return "", 0, err
	}

	return accessTokenResponse.AccessToken, accessTokenResponse.ExpiresIn, err
}
