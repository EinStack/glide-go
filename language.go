package glide

import (
	"context"
	"net/http"
)

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
	req, err := svc.client.Build(ctx, http.MethodGet, "/v1/list", nil)
	if err != nil {
		return nil, err
	}

	var resp *RouterList
	if _, err := svc.client.Send(req, resp); err != nil {
		return nil, err
	}

	return resp.Routers, nil
}

func (svc *language) Chat(ctx context.Context, req ChatRequest) (*ChatResponse, error) {
	req2, err := svc.client.Build(ctx, http.MethodPost, "/v1/chat", req)
	if err != nil {
		return nil, err
	}

	var resp *ChatResponse
	if _, err := svc.client.Send(req2, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *language) ChatStream(ctx context.Context) error {
	// TODO.
	return nil
}
