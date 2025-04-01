package gocq

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type GocqServer struct {
	conn       *websocket.Conn
	writeMutex sync.Mutex // 添加互斥锁
}

var GocqInstance *GocqServer

func NewGocqServer(conn *websocket.Conn) *GocqServer {
	return &GocqServer{
		conn: conn,
	}
}

func (g *GocqServer) SendMessage(action string, parms map[string]any) error {
	g.writeMutex.Lock()
    defer g.writeMutex.Unlock()

	messageSend := map[string]interface{}{
		"action": action,
		"params": parms,
	}
	return g.conn.WriteJSON(messageSend)
}

func (g *GocqServer) SendMessageWithEcho(action string, parms map[string]any, echo string) error {
	messageSend := map[string]interface{}{
		"action": action,
		"params": parms,
		"echo":   echo,
	}
	return g.conn.WriteJSON(messageSend)
}

func (g *GocqServer) IsConnected() bool {
	g.writeMutex.Lock()
	defer g.writeMutex.Unlock()

	// 尝试发送一个ping消息来检查连接
	err := g.conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(time.Second))
	return err == nil
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
