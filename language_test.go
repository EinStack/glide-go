package glide_test

import (
	"context"
	"testing"

	"github.com/einstack/glide-go"
)

func TestLanguage_List(t *testing.T) {
	client, _ := glide.NewClient()

	ctx := context.Background()
	if _, err := client.Language.List(ctx); err != nil {
		t.Error(err)
	}
}

func TestLanguage_Chat(t *testing.T) {
	client, _ := glide.NewClient()

	ctx := context.Background()
	req := glide.ChatRequest{}
	if _, err := client.Language.Chat(ctx, req); err != nil {
		t.Error(err)
	}
}
