package main

import (
	"context"
	"fmt"
	"log"

	ollamaembedder "github.com/henomis/lingoose/embedder/ollama"
	"github.com/henomis/lingoose/llm/ollama"
	"github.com/henomis/lingoose/thread"
)

func main() {
	myThread := thread.New().AddMessage(
		thread.NewUserMessage().AddContent(
			thread.NewTextContent("1+1"),
		),
	)
	err := ollama.New().WithEndpoint("http://localhost:11434/api").WithModel("mistral:instruct").
		WithStream(func(s string) {
			fmt.Print(s)
		}).Generate(context.Background(), myThread)
	if err != nil {
		log.Fatal(err)
	}
	embeddings, err := ollamaembedder.New().
		WithEndpoint("http://localhost:11434/api").
		WithModel("mistral").
		Embed(
			context.Background(),
			[]string{"What is the NATO purpose?"},
		)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	log.Printf("dim=%d", len(embeddings[0]))
}
