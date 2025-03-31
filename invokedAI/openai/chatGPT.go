package myopenai

import (
	"context"
	"log/slog"

	openai "github.com/sashabaranov/go-openai"
	"github.com/viogami/Gobot-vio/config"
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
	s.APIKey = config.EnvConst.ChatGPTAPIKey
	s.URL_proxy = config.EnvConst.ChatGPTURL_proxy
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

func (s *ChatGPTService) InvokeChatGPTAPI(text string) (string, error) {
	resp, err := s.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
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
		return Resp, err
	}
	return resp.Choices[0].Message.Content, err
}

func (s *ChatGPTService) InvokeChatGPTAPIWithRole(text string, role string) (string, error) {
	s.Role = role
	return s.InvokeChatGPTAPI(text)
}
