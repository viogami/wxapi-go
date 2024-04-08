# wxapi-go
**本仓库只实现了最简单的公众号服务创建，以及调用chatgpt回复。API均取自官方文档，由于需求增加，并且有现成的golang开发的微信sdk库，不再造轮子。后续所有开发转移到[wxSDK](https://github.com/viogami/wxSDK)**

**注意，业务模块中的menu.go文件中方法尚未实现,请参照[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Creating_Custom-Defined_Menu.html)**

## 介绍
这是一个使用go语言搭建的微信公众号后端，使用gin框架。根路由，get进行验证，post进行回复消息. 主要实现了一个微信公众号后端的基本功能。

完成了微信公众号的后端验证，并调用自建的chatgpt接口，可以实现公众号对用户消息基于chatgpt的回复。

主要任务为问答，均采用http通讯，实现如下：

```
    微信服务器 (WeChat)
        |
    Get | Post                    
        v   
    本地后端 (Local Backend)       HTTP
        |
        | Post
        v   
    ChatGPT后端 (ChatGPT Backend)
```

由于基于http通讯，服务端无法主动向用户端发信息，但是调用微信官方的接口可以实现这一点。首先需要在公众号后台获取开发者id和密码，使用get方法获取`access_token`。该token会定时刷新，所以要定时重新获取。

但是需要注意的是，个人微信公众号的接口权限很少，如果想主动发消息给用户一般是使用客服接口，但是个人订阅号没有客服消息的接口权限。

本仓库已经实现了对微信接口的认证操作内容(获取Access_token)，并使用post给微信后端，使公众号可以主动发起通讯。

----

但是微信公众号始终不是以聊天为目的，官方的诸多限制使得调用公众号作为调用chatgpt的媒介不是一个好的选择：
 - 微信公众号如果在用户发消息后5秒内没有发送消息，微信将会重新发起，这对于chatgpt的响应时长+网络延迟来说，5秒是不够的，特别是对文本长度有要求。

 解决方法： 在后端收到用户消息后必须回复`success`或者空字符串，才不会超时重传。避免超时无响应，可以在每次收到消息后回复自定义语句或者直接`success`，然后通过客服消息接口，主动将chatgpt的响应发送给用户。

 ----
通过搭建公众号后端，深入理解了服务器之间的通信，了解了http，ws和webhook之间的区别和联系，也对gin框架有了更深的认识。
