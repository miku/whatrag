// NAME
// all-minilm:33m-l12-v2-fp16
// all-minilm:latest
// duckdb-nsql:7b-fp16
// duckdb-nsql:latest
// falcon:7b-instruct-fp16
// gemma:7b-instruct-fp16
// gemma:7b-text-fp16
// gemma:instruct
// gemma:latest
// llama-pro:instruct
// llama2:13b-text-q6_K
// llama2:latest
// llava:13b-v1.5-q8_0
// llava:13b-v1.6-vicuna-q6_K
// llava:13b-v1.6-vicuna-q8_0
// llava:34b
// llava:34b-v1.6
// llava:34b-v1.6-q3_K_L
// llava:34b-v1.6-q4_K_S
// llava:7b-v1.6-mistral-fp16
// llava:7b-v1.6-mistral-q4_K_M
// llava:latest
// mistral:7b-instruct-v0.2-fp16
// mistral:7b-text-fp16
// mistral:instruct
// mistral:latest
// mixtral:8x7b-instruct-v0.1-q3_K_L
// nomic-embed-text:latest
// openhermes:7b-mistral-v2.5-fp16
// stable-code:3b-code-fp16
// stablelm2:1.6b-zephyr-fp16
// stablelm2:latest
// tinydolphin:1.1b-v2.8-fp16
// tinydolphin:latest
// tinyllama:1.1b-chat-v1-fp16
// tinyllama:latest
// mattw/dockerit:latest
// miku/wa-0:latest
// miku/wa-1:latest
// miku/wa-2:latest

package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/henomis/lingoose/llm/ollama"
	"github.com/henomis/lingoose/thread"
)

var model = flag.String("m", "stablelm2:1.6b-zephyr-fp16", "ollama model name")

func main() {
	flag.Parse()
	log.Printf("using default ollama endpoint with model %s", *model)
	myThread := thread.New().AddMessage(
		thread.NewUserMessage().AddContent(
			thread.NewTextContent("tell me a joke about geese"),
		),
	)
	err := ollama.New().WithModel(*model).Generate(context.Background(), myThread)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(myThread)
}
