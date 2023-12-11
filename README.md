# GoBot 

![Go Reference](https://pkg.go.dev/badge/github.com/go-telegram-bot-api/telegram-bot-api/v5.svg)
![Test](https://github.com/go-telegram-bot-api/telegram-bot-api/actions/workflows/test.yml/badge.svg)

使用go搭建的bot，取名为 **Vio** ，旨在提供一个接口，用来完成chatgpt聊天等任务，部署在一个服务器可以多个平台共同调用。

## 写在前面/preface
有部署聊天机器人的想法，但是我使用的国内服务器，而且服务器性能也堪忧，于是决定不要laas--即用云服务器部署了，找个国外的Paas平台，把写的后端送上去就好了，而且一般也都有免费计划，够用了。

## 对比表格/compare popular Paas
| 服务提供商  | Fly.io          | Railway        | Render         | Glitch         | Adaptable      | **Zeabur**      |
|-------------|-----------------|----------------|----------------|----------------|----------------|----------------|
| 长时间不活动关闭 | 否             | 否              | 15 分钟         | 5 分钟          | 是*            | 否             |
| 需要信用卡    | 是             | 是             | 否             | 否             | 否             | 否             |
| 免费计划      | 免费创建三个最低配应用| 试用$5额度| 750 小时       | 1000 小时      | 无*            | $5/mo        |
| 内存          | 256MB          | 512MB          | 512MB          | 512MB          | 256MB          | 512MB          |
| 磁盘空间      | 3GB            | 1GB            | -              | 200MB*         | 1GB            | 1GB            |
| 可写磁盘      | 是             | -              | 否             | 是             | 是*            | 是             |
| 网络带宽      | 160GB          | $0.10/GB       | 100GB          | 4000 次请求/时  | 100GB          | -              |
| 可用 Dockerfile | 是             | 是             | 是             | 否             | 否*            | 是             |
| GitHub 集成  | 是            | 是             | 是             | 是             | 是             | 是             |

对于
 - Cyclic
 - Cloudflare Workers 
 - AirCode（国内团队做的） 
 - Vercel
它们主要侧重于`Fullstack Javascript Apps - Deploy and Host in Seconds`
对非 Nodejs 的后端参考意义不大。
主要思路都是 Edge Network + Serverless Functions（函数代码在轻量级的 V8 沙盒中执行）

## 对于我/for me
只使用过cloudflare的workers部署服务，但是只能使用js，不熟悉还是挺难搞了，但是用js的可能比较舒服。
应该是完全免费的，只要绑了自己的域名在CF上，CF也提供子域名

目前感觉Zeabur挺不错，~~主要看重免费计划~~。国内社区，discord上回复也很即使，一键部署挺快的，github集成。
只部署一个机器人接口就好了，无论什么聊天平台，返回的数据基本都是互通的。

## 
表格和paas平台对比参考：[免费的 PaaS 平台汇总][1]

  [1]: https://liduos.com/Summary-of-free-PaaS-platforms.html

