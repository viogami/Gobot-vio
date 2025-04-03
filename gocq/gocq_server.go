package gocq

import (
	"sync"

	redis "github.com/redis/go-redis/v9"
)

type GocqServer struct {
	RedisClient *redis.Client
	Sender      *GocqSender
}

// 单例模式，确保只有一个GocqServer实例
var Instance *GocqServer
var once sync.Once

func NewGocqServer(redis *redis.Client) *GocqServer {
	once.Do(func() {
		Instance = &GocqServer{
			RedisClient: redis,
		}
	})
	return Instance
}
