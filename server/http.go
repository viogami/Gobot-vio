package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/viogami/Gobot-vio/chatgpt"
)

// 提取post中的msg字符串，调用chatgpt api，返回响应回答
func GptMsgHandle(w http.ResponseWriter, r *http.Request) {
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
			gptResponse = "gpt调用失败了😥 error:\n" + err.Error()
		}
		fmt.Fprintln(w, gptResponse)
	} else {
		http.Error(w, "Error: wrong HTTP method:"+r.Method+",required POST.", http.StatusMethodNotAllowed)
	}
}
