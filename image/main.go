package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/peterbourgon/ff/v3"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

const (
	envPrefix = "IMAGE_GENERATOR"
)

type Config struct {
	Prompt string
}

func main() {
	var cfg Config
	fs := flag.NewFlagSet("app", flag.ExitOnError)
	fs.StringVar(&cfg.Prompt, "prompt", "", "Image prompt")

	if err := ff.Parse(fs, os.Args[1:], ff.WithEnvVarPrefix(envPrefix)); err != nil {
		fs.Usage()
		log.Fatalf("invalid args, err: %v", err)
	}

	if cfg.Prompt == "" {
		log.Fatalf("An image generation prompt is required")
	}

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	ctx := context.Background()

	req := openai.ImageRequest{
		Model:  openai.CreateImageModelDallE3,
		N:      1,
		Prompt: cfg.Prompt,
	}

	resp, err := client.CreateImage(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Revised prompt:", resp.Data[0].RevisedPrompt)
	fmt.Println("Generated Image URL:", resp.Data[0].URL)
}
