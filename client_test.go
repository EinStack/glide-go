package glide_test

import (
	"github.com/einstack/glide-go"
	"testing"
)

func TestNewClient(t *testing.T) {
	if _, err := glide.NewClient(
		glide.WithApiKey("testing"),
		glide.WithUserAgent("Axiston/1.0"),
	); err != nil {
		t.Error(err)
	}
}

func TestClient_Health(t *testing.T) {
	client, _ := glide.NewClient(
		glide.WithApiKey("testing"),
	)

	if err := client.Health(); err != nil {
		t.Error(err)
	}
}
