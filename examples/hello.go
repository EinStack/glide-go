package main

import (
	"context"
	"log"

	"github.com/einstack/glide-go"
	"github.com/einstack/glide-go/lang"
)

var router = "myrouter"

func main() {
	client, err := glide.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	if _, err := client.Health(ctx); err != nil {
		log.Fatal(err)
	}

	req := lang.NewChatRequest("Hello")
	resp, err := client.Lang.Chat(ctx, router, req)
	if err != nil {
		log.Fatal(err)
	}

	println("response: ", resp.Content())
}
