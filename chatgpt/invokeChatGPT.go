package chatgpt

import (
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

// 定义环境变量
var (
	chatGPTAPIKey = os.Getenv("chatGPTAPIKey")
)

const (
	chatGPTURL_chat   = "https://api.openai.com/v1"
	chatGPTURL_img    = "https://api.openai.com/v1/images/generations"
	chatGPTURL_mood   = "https://api.openai.com/v1/moderations"
	chatGPTURL_broker = "https://one-api.bltcy.top/v1"
)

func InvokeChatGPTAPI(text string) (string, error) {
	config := openai.DefaultConfig(chatGPTAPIKey)
	config.BaseURL = chatGPTURL_broker

	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 1024, // 限制最大返回token，提速
			Messages: []openai.ChatCompletionMessage{
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
