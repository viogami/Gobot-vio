package chatgpt

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
	"github.com/viogami/Gobot-vio/config"
)

// 定义环境变量
var (
	chatGPTAPIKey  string
	chatGPTURL_proxy string
)

func InvokeChatGPTAPI(text string) (string, error) {
	chatGPTAPIKey = config.EnvConst.ChatGPTAPIKey
	chatGPTURL_proxy = config.EnvConst.ChatGPTURL_proxy

	prompt := GPTpreset["vio"]
	conf := openai.DefaultConfig(chatGPTAPIKey)
	conf.BaseURL = chatGPTURL_proxy

	client := openai.NewClientWithConfig(conf)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: prompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: text,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, err
}
