package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/peterbourgon/ff/v3"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

const (
	envPrefix = "CHAT"
)

type Config struct {
	SystemPrompt string
}

func main() {
	var cfg Config
	fs := flag.NewFlagSet("app", flag.ExitOnError)
	fs.StringVar(&cfg.SystemPrompt, "system", "you are a helpful chatbot", "System prompt")

	if err := ff.Parse(fs, os.Args[1:], ff.WithEnvVarPrefix(envPrefix)); err != nil {
		fs.Usage()
		log.Fatalf("invalid args, err: %v", err)
	}

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	req := openai.ChatCompletionRequest{
		Temperature: 0.01,
		MaxTokens:   2056,
		Model:       openai.GPT3Dot5Turbo16K,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: cfg.SystemPrompt,
			},
		},
	}
	fmt.Println("Conversation")
	fmt.Println("---------------------")
	fmt.Print("> ")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: s.Text(),
		})
		resp, err := client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}
		fmt.Printf("%s\n\n", resp.Choices[0].Message.Content)
		req.Messages = append(req.Messages, resp.Choices[0].Message)
		fmt.Print("> ")
	}

}
