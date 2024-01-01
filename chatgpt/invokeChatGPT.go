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
	chatGPTURL_broker = "https://one-api.bltcy.top/v1"
)

func InvokeChatGPTAPI(text string) (string, error) {
	// appConfig := flag.String("config", "./app.yaml", "application config path")
	// conf, _ := config.ConfigParse(*appConfig)
	// chatGPTAPIKey = conf.Chatgpt.chatGPTAPIKey

	config := openai.DefaultConfig(chatGPTAPIKey)
	config.BaseURL = chatGPTURL_broker

	client := openai.NewClientWithConfig(config)
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
