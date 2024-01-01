package main

import (
	"Gobot-vio/chatgpt"
	"Gobot-vio/tgbot"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var port = os.Getenv("PORT")

	// 设置 /post 路径的 HTTP 处理函数
	http.HandleFunc("/post", handlePost)

	// 启动 Web 服务器监听 port 端口
	go func() {
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			log.Println("Error starting HTTP server:", err)
		}
		log.Println("HTTP server is running on port", port)
	}()

	//创建一个tgbot
	tgbot.CreateTgbot()
}

// 提取post中的msg字符串，调用chatgpt api，返回响应回答
func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// 获取表单数据
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusInternalServerError)
			return
		}
		// 读取请求体
		postmsg := r.Form.Get("usermsg")
		if postmsg != "" {
			log.Println("POST request,the usermsg:", postmsg)
		} else {
			http.Error(w, "Error:Don`t find the key:usermsg in the POST,maybe it`s a nil", http.StatusBadRequest)
		}
		// 调用ChatGPT API
		gptResponse, err := chatgpt.InvokeChatGPTAPI(postmsg)
		if err != nil {
			log.Printf("Error calling ChatGPT API: %v", err)
			gptResponse = "gpt调用失败了😥 错误信息：\n" + err.Error()
		}
		fmt.Fprintln(w, gptResponse)
	} else {
		http.Error(w, "Error: wrong HTTP method:"+r.Method+",required POST.", http.StatusMethodNotAllowed)
	}
}
