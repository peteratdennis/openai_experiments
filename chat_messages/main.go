package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

func main() {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Temperature: 0.01,
		MaxTokens:   2056,
		N:           1,
		Model:       openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "you are a helpful chatbot",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Who won the world series in 2020?",
			},
			{
				Role:    openai.ChatMessageRoleAssistant,
				Content: "The Los Angeles Dodgers won the World Series in 2020.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Where was it played?",
			},
		},
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeText,
		},
	}

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PromptTokens:", resp.Usage.PromptTokens)
	fmt.Println("CompletionTokens:", resp.Usage.CompletionTokens)
	fmt.Println("TotalTokens:", resp.Usage.TotalTokens)
	fmt.Println("Response:", resp.Choices[0].Message.Content)
}
