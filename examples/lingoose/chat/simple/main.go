package main

import (
	"context"
	"fmt"
	"log"

	"github.com/henomis/lingoose/chat"
	"github.com/henomis/lingoose/llm/ollama"
	"github.com/henomis/lingoose/prompt"
	"github.com/henomis/lingoose/thread"
)

func main() {

	ch := chat.New(
		chat.PromptMessage{
			Type:   chat.MessageTypeSystem,
			Prompt: prompt.New("You are a professional joke writer"),
		},
		chat.PromptMessage{
			Type:   chat.MessageTypeUser,
			Prompt: prompt.New("Write a joke about a goose"),
		},
	)
	m := ollama.New().WithModel("mistral")
	messages, err := ch.ToMessages()
	if err != nil {
		log.Fatal(err)
	}
	t := thread.New()
	for _, msg := range messages {
		switch msg.Type {
		case chat.MessageTypeSystem:
			t.AddMessage(thread.NewSystemMessage().AddContent(thread.NewTextContent(msg.Content)))
		case chat.MessageTypeUser:
			t.AddMessage(thread.NewUserMessage().AddContent(thread.NewTextContent(msg.Content)))
		}
	}
	m.Generate(context.Background(), t)
	fmt.Println(t)
}
