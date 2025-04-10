package myopenai

import (
	"context"
	"log/slog"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type ChatGPTService struct {
	APIKey           string
	URL_PROXY        string
	Role             string
	Character        string
	characterSetting string

	client *openai.Client
}

func NewChatGPTService() *ChatGPTService {
	s := new(ChatGPTService)
	s.APIKey = os.Getenv("ChatGPTAPIKey")
	s.URL_PROXY = os.Getenv("ChatGPTURL_PROXY")
	s.Role = openai.ChatMessageRoleUser
	s.Character = "vio"
	s.characterSetting = gpt_preset[s.Character]

	conf := openai.DefaultConfig(s.APIKey)
	conf.BaseURL = s.URL_PROXY

	s.client = openai.NewClientWithConfig(conf)

	return s
}

func (s *ChatGPTService) SetCharacter(character string) {
	s.Character = character
	s.characterSetting = gpt_preset[character]
}

func (s *ChatGPTService) InvokeChatGPTAPI(text string) string {
	resp, err := s.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o20241120,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: s.characterSetting,
				},
				{
					Role:    s.Role,
					Content: text,
				},
			},
		},
	)
	if err != nil {
		slog.Error("Error calling ChatGPT API", "error", err)
		Resp := "AI调用失败了😥 error:\n" + err.Error()
		return Resp
	}
	return resp.Choices[0].Message.Content
}

func (s *ChatGPTService) InvokeChatGPTAPIWithRole(text string, role string) string {
	s.Role = role
	return s.InvokeChatGPTAPI(text)
}
