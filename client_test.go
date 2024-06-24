package glide_test

import (
	"context"
	"testing"

	"github.com/einstack/glide-go"
	"github.com/einstack/glide-go/config"
)

func TestNewClient(t *testing.T) {
	if _, err := glide.NewClient(
		glide.WithApiKey("testing"),
		glide.WithRawBaseURL("http://127.0.0.1:9098/"),
		glide.WithUserAgent("Einstack/1.0"),
	); err != nil {
		t.Fatal(err)
	}
}

func TestNewClientFromConfig(t *testing.T) {
	if _, err := glide.NewClientFromConfig(
		config.NewConfig(),
	); err != nil {
		t.Fatal(err)
	}
}

func TestClient_Health(t *testing.T) {
	client, _ := glide.NewClient(
		glide.WithApiKey("testing"),
	)

	ctx := context.Background()
	if _, err := client.Health(ctx); err != nil {
		t.Fatal(err)
	}
}
