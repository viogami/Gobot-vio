# GoBot 

![Go Reference](https://pkg.go.dev/badge/github.com/go-telegram-bot-api/telegram-bot-api/v5.svg)

使用go搭建的bot，取名为 **Vio** ，旨在提供一个接口，用来接受不同协议的请求，调用转发外部讯息并返回，目前可以完成基于chatgpt聊天任务，部署在一个服务器可以多个平台共同调用。

支持通信方式：
 - http
 - webhook
 - 反向websocket

实现平台:
 - [x] Telegram bot
 - [x] QQ bot
 - [ ] 微信bot

## 写在前面/preface
有部署聊天机器人的想法，但是我使用的国内服务器，而且服务器性能也堪忧，于是决定不要laas--即用云服务器部署了，找个国外的Paas平台，把写的后端送上去就好了，而且一般也都有免费计划，够用了。

传统的聊天机器人服务都是一体化的，和聊天平台需要集成。我希望把消息处理的逻辑和平台部署的逻辑做两个服务，后者发送信息给前者，前者返回需要发送的信息，后者再在聊天平台呈现信息。

目前已经完成了：

http：
  - 该后端天然支持http请求，使用go原生net包，创建了一个post请求的路由，可以解析post内容转发调用chatgpt，目前我将其使用在微信公众号的后端上。

webhook
 - 为Tg设置了webhook，可以监听tg服务器的消息，实现tgbot。
 - 针对Telegram的消息处理，对私人，群组，超级群组各有不同的应答模式。

 反向ws：
  - 配合[go-cqhttp](https://github.com/Mrs4s/go-cqhttp)使用，用于实现qqbot。

chatgpt调用：
 - 在聊天中支持chatgpt的回复。该功能使用了[go-openai](https://github.com/sashabaranov/go-openai)库，支持了go语言对openai api的调用


## Paas平台推荐

目前我正在使用heroku这个paas平台，国外的，可以通过github学生包免费使用。<br>
之前使用过Zeabur这个平台，国人开发的，discord上官方也很活跃，如果服务请求不多，可以使用，每月有5刀的额度。

对于
 - Cloudflare Workers 
 - AirCode（国内团队做的） 
它们主要侧重于
>Fullstack Javascript Apps - Deploy and Host in Seconds

对非 Nodejs 的后端参考意义不大。
主要思路都是 Edge Network + Serverless Functions（函数代码在轻量级的 V8 沙盒中执行）
 - Fly.io
 - Railway

 这两需要绑信用卡，并且railway免费计划只有试用5$了


 看重免费计划可以查询 [Free for dev](https://github.com/ripienaar/free-for-dev)中的列表.

### 对于我/for me

使用过cloudflare的workers部署服务，只能使用js，不熟悉还是挺难搞了。是完全免费的，只要绑了自己的域名在CF上，CF也提供子域名

Zeabur挺不错，国内社区，discord上回复也很即使，一键部署挺快的，github集成。
zeabur上项目部署非常快,甚至不用写dockfile,而且对go项目有完整的支持,算是符合他们的口号:
> Deploying your service with one click

但是Zeabur每月5$不够用了，现在我用github学生包的heroku平台，注册需要关闭adguard，需要绑定国外银行卡，绑卡时建议关闭梯子。

只部署一个机器人接口就好了.无论什么聊天平台，通讯功能的实现基本都是互通的。

本后端最终希望实现只对外暴露一个API,实现机器人通讯的应答模式,对不同平台创建不同的新服务,调用接口皆可进行通讯服务.

## Paas部署注意点
### 端口号
注意一下项目的端口号设置,最好设置在环境变量中,然后在项目中通过`os.Getenv("xxx")`来获取端口号.

zeabur的go项目中，环境变量`PORT`是默认8080，且为全局的。也可以不设置，直接调用就好了。

 - 对于tgbot：官方示例中使用的8443端口，在部署到paas平台时8443端口需要确认是否开放。
  我建议不要使用官方示例中把端口号写明的写法，通过环境变量`PORT`调用端口号，避免webhook创建失败，或者监听未开放的端口等问题。

### 证书问题
zeabur部署项目自带证书,做完域名映射可以直接https访问.
所以在设置webhook进行和tg服务器通讯的时候不需要手动加载`cert.pem`和`key.pem`

在部署tg的bot时,可以修改tgbot官方对go语言搭建bot示例中的:
``` go
  ...

  log.Printf("Authorized on account %s", bot.Self.UserName)
  wh, _ := tgbotapi.NewWebhook(TG_WEBHOOK_URL + bot.Token)

  ...

  go http.ListenAndServe(":"+port, nil)
```
直接使用`NewWebhook`和`ListenAndServe`函数即可.

### 环境变量
在目前我的实现中，定义了三个环境变量
```env
TG_WEBHOOK_URL=https://yousite.com/tgbot
BOT_TOKEN=your token
chatGPTAPIKey=sk-your key
```

