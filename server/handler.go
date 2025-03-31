package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/viogami/Gobot-vio/gocq"
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
			log.Println("POST request,the usermsg:", postmsg)
		} else {
			http.Error(w, "Error:Don`t find the key:usermsg in the POST,maybe it`s a nil", http.StatusBadRequest)
		}
		// 调用ChatGPT API
		gptResponse, err := NewAIServer().ProcessMessage(postmsg)
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

	for {
		// 从WebSocket连接读取消息
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		// 打印接收到的消息
		err = gocq.Log_post_type(p)
		if err != nil {
			log.Println(err)
		} else {
			// 发送消息
			message_send := gocq.Handle_event(p)
			if len(message_send) != 0 {
				go func() {
					for _, msg := range message_send {
						// 发送消息
						err = conn.WriteJSON(msg)
						if err != nil {
							log.Println(err)
							return
						}
					}
				}()
			}
		}
	}
}
