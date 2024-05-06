package glide

import (
	"context"
)

// RouterConfig TODO.
type RouterConfig struct {
}

// ChatRequest TODO.
type ChatRequest struct {
}

// NewChatRequest instantiates a new ChatRequest.
func NewChatRequest() ChatRequest {
	// TODO.
	return ChatRequest{}
}

// ChatResponse TODO.
type ChatResponse struct {
}

// LanguageSvc implements APIs for '/v1/language' endpoints.
type LanguageSvc interface {
	// List retrieves a list of all router configs.
	List(ctx context.Context) ([]RouterConfig, error)
	// Chat sends a single chat request to a specified router and retrieves the response.
	Chat(ctx context.Context, req ChatRequest) (*ChatResponse, error)
	// ChatStream establishes a WebSocket connection for streaming chat messages from a specified router.
	ChatStream(ctx context.Context) error
}

type language struct {
	client *Client
}

func (svc *language) List(ctx context.Context) ([]RouterConfig, error) {
	// TODO.
	return nil, nil
}

func (svc *language) Chat(ctx context.Context, req ChatRequest) (*ChatResponse, error) {
	// TODO.
	return nil, nil
}

func (svc *language) ChatStream(ctx context.Context) error {
	// TODO.
	return nil
}
