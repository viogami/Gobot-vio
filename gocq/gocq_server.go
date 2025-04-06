package gocq

import (
	"sync"

	redis "github.com/redis/go-redis/v9"
)

type GocqServer struct {
	RedisClient *redis.Client
	Sender      *GocqSender
	MsgQueue    chan []byte // 消息队列
	ResponseMap sync.Map    // 用于存储等待响应的 echo 和 channel
}

// 单例模式，确保只有一个GocqServer实例
var Instance *GocqServer
var once sync.Once

func NewGocqServer(redis *redis.Client) *GocqServer {
	once.Do(func() {
		Instance = &GocqServer{
			RedisClient: redis,
			MsgQueue:    make(chan []byte, 100), // 初始化消息队列，大小为100
			ResponseMap: sync.Map{},             // 初始化响应映射
		}
	})
	return Instance
}
