package main

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type ChatGPTRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

func invokeChatGPTAPI(text string) (string, error) {
	client := resty.New()

	requestData := ChatGPTRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: text},
		},
		MaxTokens: 256,
	}

	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+chatGPTAPIKey).
		SetBody(requestData).
		Post(chatGPTAPIURL)

	if err != nil {
		return "", err
	}

	var responseData ChatGPTResponse
	err = json.Unmarshal(response.Body(), &responseData)
	if err != nil {
		return "", err
	}

	return responseData.Choices[0].Message.Content, nil
}
