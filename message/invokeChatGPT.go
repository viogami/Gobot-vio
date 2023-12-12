package message

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
	chatGPTURL_chat   = "https://api.openai.com/v1/chat/completions"
	chatGPTURL_img    = "https://api.openai.com/v1/images/generations"
	chatGPTURL_mood   = "https://api.openai.com/v1/moderations"
	chatGPTURL_broker = "https://one-api.bltcy.top"
)

func invokeChatGPTAPI(text string) (string, error) {
	config := openai.DefaultConfig(chatGPTAPIKey)
	config.BaseURL = chatGPTURL_broker
	client := openai.NewClient(chatGPTAPIKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
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
