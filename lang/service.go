package lang

import (
	"context"
	"fmt"
	"net/http"

	"github.com/einstack/glide-go/config"
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
	config *config.Config
}

// NewLanguage instantiates a new Language service.
func NewLanguage(config *config.Config) Language {
	return &languageSvc{config: config}
}

func (svc *languageSvc) List(ctx context.Context) (*RouterList, error) {
	httpReq, err := svc.config.Build(ctx, http.MethodGet, "/v1/list", nil)
	if err != nil {
		return nil, err
	}

	var resp *RouterList
	if _, err := svc.config.Send(httpReq, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *languageSvc) Chat(ctx context.Context, router string, req ChatRequest) (*ChatResponse, error) {
	path := fmt.Sprintf("/v1/%s/chat", router)
	httpReq, err := svc.config.Build(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, err
	}

	var resp *ChatResponse
	if _, err := svc.config.Send(httpReq, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *languageSvc) ChatStream(ctx context.Context, router string) (Chat, error) {
	path := fmt.Sprintf("/v1/%s/chatStream", router)
	conn, err := svc.config.Upgrade(ctx, path)
	if err != nil {
		return nil, err
	}

	return newChatService(conn), nil
}
