package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/henomis/lingoose/document"
	ollamaembedder "github.com/henomis/lingoose/embedder/ollama"
	"github.com/henomis/lingoose/index"
	"github.com/henomis/lingoose/index/vectordb/jsondb"
	"github.com/henomis/lingoose/rag"
	"github.com/henomis/lingoose/types"
)

var example = `
	Vector Database Creation: RAG starts by converting an internal dataset into
	vectors and storing them in a vector database (or a database of your
	choosing).


	User Input: A user provides a query in natural language, seeking an answer
	or completion.


	Information Retrieval: The retrieval mechanism scans the vector database to
	identify segments that are semantically similar to the user’s query (which
	is also embedded). These segments are then given to the LLM to enrich its
	context for generating responses.


	Combining Data: The chosen data segments from the database are combined
	with the user’s initial query, creating an expanded prompt.


	Generating Text: The enlarged prompt, filled with added context, is then
	given to the LLM, which crafts the final, context-aware response.
`

func main() {
	re := regexp.MustCompile(`\n\n`)
	pieces := re.Split(example, -1)
	var docs []document.Document
	for _, p := range pieces {
		p = strings.TrimSpace(p)
		if len(p) < 120 {
			continue
		}
		doc := document.Document{
			Content:  p,
			Metadata: types.Meta{},
		}
		docs = append(docs, doc)
	}
	rag := rag.New(
		index.New(
			jsondb.New().WithPersist("index.json"),
			ollamaembedder.New().WithModel("nomic-embed-text"),
		),
	).WithChunkSize(1000).WithChunkOverlap(0)
	// remove file to reset
	if _, err := os.Stat("index.json"); os.IsNotExist(err) {
		log.Printf("adding %d docs", len(docs))
		rag.AddDocuments(
			context.Background(),
			docs...,
		)
	}
	q := "user interaction"
	log.Println(q)
	results, err := rag.Retrieve(context.Background(), q)
	if err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}
