package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	ollamaembedder "github.com/henomis/lingoose/embedder/ollama"
	"github.com/henomis/lingoose/index"
	indexoption "github.com/henomis/lingoose/index/option"
	"github.com/henomis/lingoose/thread"

	"github.com/henomis/lingoose/index/vectordb/jsondb"
	"github.com/henomis/lingoose/llm/ollama"
	"github.com/henomis/lingoose/loader"
	"github.com/henomis/lingoose/prompt"
	"github.com/henomis/lingoose/textsplitter"
	"github.com/henomis/lingoose/types"
)

const (
	KB = "https://en.wikipedia.org/wiki/World_War_II"
)

func main() {

	index := index.New(
		jsondb.New().WithPersist("db.json"),
		ollamaembedder.New().WithModel("llama2"),
	).WithIncludeContents(true)

	indexIsEmpty, _ := index.IsEmpty(context.Background())

	if indexIsEmpty {
		err := ingestData(index)
		if err != nil {
			panic(err)
		}
	}

	model := ollama.New().WithModel("llama2")

	fmt.Println("Enter a query to search the knowledge base. Type 'quit' to exit.")
	query := ""
	for query != "quit" {
		fmt.Printf("> ")
		reader := bufio.NewReader(os.Stdin)
		query, _ := reader.ReadString('\n')
		if query == "quit" {
			break
		}
		similarities, err := index.Query(context.Background(), query, indexoption.WithTopK(3))
		if err != nil {
			panic(err)
		}
		content := ""
		for _, similarity := range similarities {
			fmt.Printf("Similarity: %f\n", similarity.Score)
			fmt.Printf("Document: %s\n", similarity.Content())
			fmt.Println("Metadata: ", similarity.Metadata)
			fmt.Println("----------")
			content += similarity.Content() + "\n"
		}
		systemPrompt := prompt.New("You are an helpful assistant. Answer to the questions using only " +
			"the provided context. Don't add any information that is not in the context. " +
			"If you don't know the answer, just say 'I don't know'.",
		)
		userPrompt := prompt.NewPromptTemplate(
			"Based on the following context answer to the question.\n\nContext:\n{{.context}}\n\nQuestion: {{.query}}").WithInputs(
			types.M{
				"query":   query,
				"context": content,
			},
		)
		myThread := thread.New().AddMessages(
			thread.NewSystemMessage().AddContent(
				thread.NewTextContent(systemPrompt.String()),
			),
			thread.NewUserMessage().AddContent(
				thread.NewTextContent(userPrompt.String()),
			),
		)
		err = model.Generate(context.Background(), myThread)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(myThread)
	}
}

func ingestData(index *index.Index) error {

	fmt.Printf("Learning Knowledge Base...")

	loader := loader.NewPDFToTextLoader("./kb").WithPDFToTextPath("/usr/bin/pdftotext")

	documents, err := loader.Load(context.Background())
	if err != nil {
		return err
	}

	textSplitter := textsplitter.NewRecursiveCharacterTextSplitter(2000, 200)

	documentChunks := textSplitter.SplitDocuments(documents)

	err = index.LoadFromDocuments(context.Background(), documentChunks)
	if err != nil {
		return err
	}

	fmt.Printf("Done\n")

	return nil
}
