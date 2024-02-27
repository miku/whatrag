package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/henomis/lingoose/assistant"
	ollamaembedder "github.com/henomis/lingoose/embedder/ollama"
	"github.com/henomis/lingoose/index"
	"github.com/henomis/lingoose/index/vectordb/jsondb"
	"github.com/henomis/lingoose/llm/ollama"
	"github.com/henomis/lingoose/rag"
	"github.com/henomis/lingoose/thread"
)

// download https://raw.githubusercontent.com/hwchase17/chat-your-data/master/state_of_the_union.txt

func main() {
	m := ollama.New().WithModel("mistral")
	r := rag.NewFusion(
		index.New(
			jsondb.New().WithPersist("db.json"),
			ollamaembedder.New(),
		), m,
	).WithTopK(3)
	_, err := os.Stat("db.json")
	if os.IsNotExist(err) {
		err = r.AddSources(context.Background(), "state_of_the_union.txt")
		if err != nil {
			log.Fatal(err)
		}
	}
	a := assistant.New(m).WithRAG(r).WithThread(
		thread.New().AddMessages(
			thread.NewUserMessage().AddContent(
				thread.NewTextContent("what is the purpose of NATO?"),
			),
		),
	)
	err = a.Run(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("----")
	fmt.Println(a.Thread())
	fmt.Println("----")
}
