package lang

import (
	"context"
	"io"

	"github.com/gorilla/websocket"
)

// Chat is a streaming (`WebSocket`) chat connection.
type Chat interface {
	io.Closer

	// Send attempts to send the provided chat request.
	Send(ctx context.Context) error

	// Recv attempts to receive the next chat response.
	Recv(ctx context.Context) error
}

type chatService struct {
	conn *websocket.Conn
}

// newChatService instantiates a new chatService.
func newChatService(conn *websocket.Conn) *chatService {
	return &chatService{conn: conn}
}

func (svc *chatService) Send(ctx context.Context) error {
	// TODO.
	panic("implement me")
}

func (svc *chatService) Recv(ctx context.Context) error {
	// TODO.
	panic("implement me")
}

// Close closes the underlying connection without sending or waiting for a close message.
func (svc *chatService) Close() error {
	return svc.conn.Close()
}
