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
	chatGPTAURL_chat = "https://api.openai.com/v1/chat/completions"
	chatGPTAURL_img  = "https://api.openai.com/v1/images/generations"
	chatGPTAURL_mood = "https://api.openai.com/v1/moderations"
)

func invokeChatGPTAPI(text string) (string, error) {
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
