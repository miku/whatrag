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
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/henomis/lingoose/llm/ollama"
	"github.com/henomis/lingoose/thread"
)

var (
	model = flag.String("m", "llava:34b-v1.6-q4_K_S", "ollama model name")
	image = flag.String("i", "https://golangleipzig.space/images/gridfuse-gophers-s.png", "image url or path")
)

func main() {
	flag.Parse()
	var target string = *image
	if strings.HasPrefix(*image, "http") {
		ext, err := getFileExtensionFromUrl(*image)
		if err != nil {
			ext = "png"
			log.Println("warning: falling back to png")
		}
		target = path.Join(os.TempDir(), fmt.Sprintf("imgtotext-%s.%s", generateSHA1(*image), ext))
		if _, err := os.Stat(target); os.IsNotExist(err) {
			err := downloadFile(*image, target)
			if err != nil {
				log.Fatalf("could not save file: %v", err)
			}
		}
	}
	myThread := thread.New().AddMessage(
		thread.NewUserMessage().AddContent(
			thread.NewTextContent(target),
		),
	)
	err := ollama.New().WithModel(*model).Generate(context.Background(), myThread)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(myThread)
}

func generateSHA1(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}

// downloadFile will download a url and store it in local filepath.
// It writes to the destination file as it downloads it, without
// loading the entire file into memory.
func downloadFile(url string, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func getFileExtensionFromUrl(link string) (string, error) {
	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	pos := strings.LastIndex(u.Path, ".")
	if pos == -1 {
		return "", errors.New("couldn't find a period to indicate a file extension")
	}
	return u.Path[pos+1 : len(u.Path)], nil
}
