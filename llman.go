package main

import (
	"context"
	"fmt"
	"log"

	"github.com/henomis/lingoose/llm/ollama"
	"github.com/henomis/lingoose/llm/openai"
	"github.com/henomis/lingoose/thread"
)

func main() {
	myThread := thread.New().AddMessage(
		thread.NewUserMessage().AddContent(
			thread.NewTextContent("Tell me a joke about geese"),
		),
	)
	err := ollama.New().WithEndpoint("http://localhost:11434/api").WithModel("mistral").
		WithStream(func(s string) {
		}).Generate(context.Background(), myThread)
	err := openai.New().Generate(context.Background(), myThread)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(myThread)
}
