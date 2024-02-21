package main

import (
	"context"
	"fmt"
	"log"

	"github.com/henomis/lingoose/llm/ollama"
	"github.com/henomis/lingoose/thread"
)

func main() {
	myThread := thread.New().AddMessage(
		thread.NewUserMessage().AddContent(
			thread.NewTextContent("hello world in haskell"),
		),
	)
	err := ollama.New().WithEndpoint("http://localhost:11434/api").WithModel("mistral:instruct").
		WithStream(func(s string) {
			fmt.Print(s)
		}).Generate(context.Background(), myThread)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(myThread)
}
