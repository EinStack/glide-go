package glide_test

import (
	"context"
	"github.com/einstack/glide-go"
	"testing"
)

var router = "myrouter"

func TestLanguage_List(t *testing.T) {
	client, _ := glide.NewClient()
	ctx := context.Background()

	if _, err := client.Lang.List(ctx); err != nil {
		t.Error(err)
	}
}

func TestLanguage_Chat(t *testing.T) {
	client, _ := glide.NewClient()
	ctx := context.Background()

	req := glide.NewChatRequest("Hello")
	if _, err := client.Lang.Chat(ctx, router, req); err != nil {
		t.Error(err)
	}
}

func TestLanguage_ChatStream(t *testing.T) {
	client, _ := glide.NewClient()
	ctx := context.Background()

	if _, err := client.Lang.ChatStream(ctx, router); err != nil {
		t.Error(err)
	}
}
