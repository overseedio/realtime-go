package realtimego

import (
	"context"
	"time"
)

// Socket manages a connection to the phoenix server via websockets.
type Socket struct{}

// connect creates a connection to the server.
func (s *Socket) connect(ctx context.Context, addr string) error {
	return nil
}

// disconnect closes and cleans up the connection.
func (s *Socket) disconnect() error {
	return nil
}

// push sends data on the connection.
func (s *Socket) push(data interface{}) error {
	return nil
}

// heartbeat is a routine informing the connection the client is still alive.
func (s *Socket) heartbeat(ctx context.Context, interval time.Duration) {}

// listen is a routine receiving messages from the connection.
func (s *Socket) listen(ctx context.Context) {}
