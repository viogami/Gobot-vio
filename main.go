package main

import (
	"Gobot-vio/chatgpt"
	"Gobot-vio/tgbot"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	var port = os.Getenv("PORT")
	//创建一个tgbot
	tgbot.CreateTgbot()

	// 创建一个默认的 Gin 引擎
	router := gin.Default()

	// 设置一个 POST 请求的路由
	router.POST("/post", func(c *gin.Context) {
		// 从请求中获取字符串参数
		receivemsgText := c.PostForm("receivemsg")

		// 调用ChatGPT API
		gptResponse, err := chatgpt.InvokeChatGPTAPI(receivemsgText)
		if err != nil {
			log.Printf("Error calling ChatGPT API: %v", err)
			gptResponse = "gpt调用失败了😥 错误信息：\n" + err.Error()
		}

		// 返回响应
		c.JSON(http.StatusOK, gin.H{"response": gptResponse})
	})

	// 启动 Web 服务器监听 port 端口
	err := router.Run(":" + port)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	log.Println("gin启动完毕,正在监听 ", port, "端口的post请求")
}
