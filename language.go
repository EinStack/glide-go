package glide

import (
	"context"
	"fmt"
	"net/http"
)

// Language implements APIs for '/v1/language' endpoints.
type Language interface {
	// List retrieves a list of all router configs.
	List(ctx context.Context) (*RouterList, error)
	// Chat sends a single chat request to a specified router and retrieves the response.
	Chat(ctx context.Context, router string, req ChatRequest) (*ChatResponse, error)
	// ChatStream establishes a WebSocket connection for streaming chat messages from a specified router.
	ChatStream(ctx context.Context, router string) (Chat, error)
}

type languageSvc struct {
	client *Client
}

func (svc *languageSvc) List(ctx context.Context) (*RouterList, error) {
	httpReq, err := svc.client.Build(ctx, http.MethodGet, "/v1/list", nil)
	if err != nil {
		return nil, err
	}

	var resp *RouterList
	if _, err := svc.client.Send(httpReq, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *languageSvc) Chat(ctx context.Context, router string, req ChatRequest) (*ChatResponse, error) {
	path := fmt.Sprintf("/v1/%s/chat", router)
	httpReq, err := svc.client.Build(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, err
	}

	var resp *ChatResponse
	if _, err := svc.client.Send(httpReq, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *languageSvc) ChatStream(ctx context.Context, router string) (Chat, error) {
	path := fmt.Sprintf("/v1/%s/chatStream", router)
	conn, err := svc.client.Upgrade(ctx, path)
	if err != nil {
		return nil, err
	}

	return newChatService(conn), nil
}
