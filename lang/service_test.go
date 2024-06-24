package lang_test

import (
	"context"
	"testing"

	"github.com/einstack/glide-go"
	"github.com/einstack/glide-go/lang"
)

var router = "myrouter"

func TestLanguage_List(t *testing.T) {
	client, _ := glide.NewClient()
	ctx := context.Background()

	if _, err := client.Lang.List(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestLanguage_Chat(t *testing.T) {
	client, _ := glide.NewClient()
	ctx := context.Background()

	req := lang.NewChatRequest("Hello")
	if _, err := client.Lang.Chat(ctx, router, req); err != nil {
		t.Fatal(err)
	}
}

func TestLanguage_ChatStream(t *testing.T) {
	client, _ := glide.NewClient()
	ctx := context.Background()

	chat, err := client.Lang.ChatStream(ctx, router)
	if err != nil {
		t.Fatal(err)
	}

	if err := chat.Close(); err != nil {
		t.Fatal(err)
	}
}
