// NAME
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
//
// This needs to run on the same machine where the local model runs.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/henomis/lingoose/llm/ollama"
	"github.com/henomis/lingoose/thread"
)

var (
	model = flag.String("m", "llava:34b-v1.6-q4_K_S", "ollama model name")
	image = flag.String("i", "https://golangleipzig.space/images/gridfuse-gophers-s.png", "image url or path")
)

func main() {
	flag.Parse()
	myThread := thread.New().AddMessage(
		thread.NewUserMessage().AddContent(
			thread.NewImageContentFromURL(*image),
		),
	)
	err := ollama.New().WithModel(*model).Generate(context.Background(), myThread)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(myThread)
}
