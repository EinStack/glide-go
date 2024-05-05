package glide_test

import (
	"context"
	"testing"

	"github.com/einstack/glide-go"
)

func TestNewClient(t *testing.T) {
	if _, err := glide.NewClient(
		glide.WithApiKey("testing"),
		glide.WithUserAgent("Einstack/1.0"),
	); err != nil {
		t.Error(err)
	}
}

func TestClient_Health(t *testing.T) {
	client, _ := glide.NewClient(
		glide.WithApiKey("testing"),
	)

	ctx := context.Background()
	if err := client.Health(ctx); err != nil {
		t.Error(err)
	}
}
