package realtimego

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// socket manages a connection to the phoenix server via websockets.
type socket struct {
	socket *websocket.Conn
}

func newSocket() *socket {
	return &socket{}
}

// connect creates a connection to the server.
func (s *socket) connect(ctx context.Context, addr string) error {
	conn, resp, err := websocket.DefaultDialer.Dial(addr, http.Header{})
	log.Println("connected to server:", resp)
	if err != nil {
		return err
	}
	s.socket = conn

	return nil
}

// disconnect closes and cleans up the connection.
func (s *socket) disconnect() error {
	return nil
}

// push sends data on the connection.
func (s *socket) push(data interface{}) error {
	return nil
}

// heartbeat is a routine informing the connection the client is still alive.
func (s *socket) heartbeat(ctx context.Context, interval time.Duration) {}

// listen is a routine receiving messages from the connection.
func (s *socket) listen(ctx context.Context) {}
