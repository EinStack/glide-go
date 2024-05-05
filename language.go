package glide

import "context"

// RouterConfig TODO.
type RouterConfig struct {
}

// ChatRequest TODO.
type ChatRequest struct {
}

// ChatResponse TODO.
type ChatResponse struct {
}

// LanguageSvc implements APIs for '/v1/language' endpoints.
type LanguageSvc interface {
	List(ctx context.Context) ([]RouterConfig, error)
	Chat(ctx context.Context, req ChatRequest) (*ChatResponse, error)
	// TODO. ChatStream(ctx context.Context) (<-chan ChatResponse, error)
}

type language struct {
	client *Client
}

func (impl *language) List(ctx context.Context) ([]RouterConfig, error) {
	// TODO.
	return nil, nil
}

func (impl *language) Chat(ctx context.Context, req ChatRequest) (*ChatResponse, error) {
	// TODO.
	return nil, nil
}
