package myopenai

import (
	"context"
	"log/slog"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type ChatGPTService struct {
	APIKey           string
	URL_proxy        string
	Role             string
	Character        string
	characterSetting string

	client *openai.Client
}

func NewChatGPTService() *ChatGPTService {
	s := new(ChatGPTService)
	s.APIKey = os.Getenv("ChatGPTAPIKey")
	s.URL_proxy = os.Getenv("ChatGPTURL_proxy")
	s.Role = openai.ChatMessageRoleUser
	s.Character = "vio"
	s.characterSetting = GPTpreset[s.Character]

	conf := openai.DefaultConfig(s.APIKey)
	conf.BaseURL = s.URL_proxy

	s.client = openai.NewClientWithConfig(conf)

	return s
}

func (s *ChatGPTService) SetCharacter(character string) {
	s.Character = character
	s.characterSetting = GPTpreset[character]
}

func (s *ChatGPTService) InvokeChatGPTAPI(text string) string {
	resp, err := s.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
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
		Resp := "gpt调用失败了😥 error:\n" + err.Error()
		return Resp
	}
	return resp.Choices[0].Message.Content
}

func (s *ChatGPTService) InvokeChatGPTAPIWithRole(text string, role string) string {
	s.Role = role
	return s.InvokeChatGPTAPI(text)
}
