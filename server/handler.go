package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/viogami/Gobot-vio/AI"
	"github.com/viogami/Gobot-vio/gocq"
	"github.com/viogami/Gobot-vio/gocq/event"
)

// gptMsgHandle 处理POST请求
func gptMsgHandle(w http.ResponseWriter, r *http.Request) {
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
			slog.Info("POST request,the usermsg:", "postmsg", postmsg)
		} else {
			http.Error(w, "Error:Don`t find the key:usermsg in the POST,maybe it`s a nil", http.StatusBadRequest)
		}
		// 调用ChatGPT API
		gptResponse := AI.NewAIServer().ProcessMessage(postmsg)
		fmt.Fprintln(w, gptResponse)
	} else {
		http.Error(w, "Error: wrong HTTP method:"+r.Method+",required POST.", http.StatusMethodNotAllowed)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	gocq.Instance.Sender = gocq.NewGocqSender(conn)

	for {
		// 从WebSocket连接读取消息
		_, p, err := conn.ReadMessage()
		if err != nil {
			slog.Error("Error reading message from WebSocket:", "error", err)
			return
		}
		// 处理接收到的事件
		if event.IsEvent(p) {
			e, _ := event.ParseEvent(p)
			go e.LogInfo()
			go e.Handle()
			continue
		}
	}
}
