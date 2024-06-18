package glide_test

import (
	"context"
	"testing"

	"github.com/einstack/glide-go"
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

	req := glide.NewChatRequest()
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
