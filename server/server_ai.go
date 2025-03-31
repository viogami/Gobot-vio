package server

import (
	openai "github.com/viogami/Gobot-vio/invokedAI/openai"
)

type AIServer struct {
	GptServer *openai.ChatGPTService
}

func(s *AIServer) ProcessMessage(message string) (string, error) {
    return s.GptServer.InvokeChatGPTAPI(message)
}

func NewAIServer() *AIServer {
	s := new(AIServer)
	s.GptServer = openai.NewChatGPTService()
	return s
}