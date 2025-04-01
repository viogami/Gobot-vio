package gocq

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type GocqServer struct {
	conn       *websocket.Conn
	writeMutex sync.Mutex // 添加互斥锁，ws无并发安全
}

var GocqInstance *GocqServer

func NewGocqServer(conn *websocket.Conn) *GocqServer {
	return &GocqServer{
		conn:       conn,
		writeMutex: sync.Mutex{}, // 初始化互斥锁
	}
}

func (g *GocqServer) SendToGocq(action string, params map[string]any) error {
	g.writeMutex.Lock()
	defer g.writeMutex.Unlock()

	messageSend := map[string]interface{}{
		"action": action,
		"params": params,
	}

	err := g.conn.WriteJSON(messageSend)
	if err != nil {
		return err
	}
	// 等待响应
	_, msg, err := g.conn.ReadMessage()
	if err != nil {
		return err
	}
	// 解析消息
	var response GocqResp
	err = json.Unmarshal(msg, &response)
	if err != nil {
		return err
	}
	if response.Status != "ok" {
		return fmt.Errorf("error: %s, msg: %s,wording: %s", response.Status, response.Msg, response.Wording)
	}
	return nil
}

func (g *GocqServer) SendMessageWithEcho(action string, params map[string]any, echo string) error {
	messageSend := map[string]interface{}{
		"action": action,
		"params": params,
		"echo":   echo,
	}
	return g.conn.WriteJSON(messageSend)
}

func (g *GocqServer) Reconnect(url string) error {
	g.writeMutex.Lock()
	defer g.writeMutex.Unlock()
	// 关闭旧连接
	if g.conn != nil {
		g.conn.Close()
	}
	// 建立新连接
	newConn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}
	g.conn = newConn
	return nil
}
