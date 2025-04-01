package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/viogami/Gobot-vio/AIServer"
	"github.com/viogami/Gobot-vio/gocq"
	"github.com/viogami/Gobot-vio/gocq/event"
)

// GptMsgHandle 处理POST请求
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
			slog.Info("POST request,the usermsg:", "postmsg", postmsg)
		} else {
			http.Error(w, "Error:Don`t find the key:usermsg in the POST,maybe it`s a nil", http.StatusBadRequest)
		}
		// 调用ChatGPT API
		gptResponse := AIServer.NewAIServer().ProcessMessage(postmsg)
		fmt.Fprintln(w, gptResponse)
	} else {
		http.Error(w, "Error: wrong HTTP method:"+r.Method+",required POST.", http.StatusMethodNotAllowed)
	}
}

// GocqWsHandle 处理WebSocket请求
func GocqWsHandle(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// 创建一个GocqServer单例
	gocq.GocqInstance = gocq.NewGocqServer(conn)
	for {
		// 从WebSocket连接读取消息
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		
		// 处理接收到的消息
		e, err := event.ParseEvent(p)
		if err != nil {
			slog.Error("Error: ", "err", err)
			continue
		}
		go e.Slog()
		go e.Do()

		// 主动发送消息

		// // 打印接收到的消息
		// err = gocq.Log_post_type(p)
		// if err != nil {
		// 	slog.Error("Error: ","err",err)
		// } else {
		// 	// 发送消息
		// 	message_send := gocq.Handle_event(p)
		// 	if len(message_send) != 0 {
		// 		go func() {
		// 			for _, msg := range message_send {
		// 				// 发送消息
		// 				err = conn.WriteJSON(msg)
		// 				if err != nil {
		// 					slog.Error("Error: ", "err", err)
		// 					return
		// 				}
		// 			}
		// 		}()
		// 	}
		// }
	}
}
