package glide

import (
	"context"
	"github.com/gorilla/websocket"
	"io"
)

// Chat is a streaming (`WebSocket`) chat connection.
type Chat interface {
	io.Closer

	// Send TODO.
	Send(ctx context.Context) error

	// Recv TODO.
	Recv(ctx context.Context) error
}

type chatService struct {
	conn *websocket.Conn
}

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
