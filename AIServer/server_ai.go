package AIServer

import (
	openai "github.com/viogami/Gobot-vio/AIServer/openai"
)

type AIServer struct {
	GptServer *openai.ChatGPTService
}

func (s *AIServer) ProcessMessage(message string) string {
	return s.GptServer.InvokeChatGPTAPI(message)
}

func NewAIServer() *AIServer {
	s := new(AIServer)
	s.GptServer = openai.NewChatGPTService()
	return s
}
