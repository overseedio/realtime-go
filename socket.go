package realtimego

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// socket manages a connection to the phoenix server via websockets.
type socket struct {
	socket *websocket.Conn
	router *router
	cancel context.CancelFunc
	mu     sync.Mutex

	// heartbeatInterval is the delay in seconds between heartbeat notifications to the server.
	heartbeatInterval uint
}

func newSocket(heartbeatInterval uint) *socket {
	// default interval
	if heartbeatInterval == 0 {
		heartbeatInterval = 10
	}

	return &socket{heartbeatInterval: heartbeatInterval}
}

// connect creates a connection to the server.
func (s *socket) connect(ctx context.Context, addr string) error {
	// socket context
	ctx, cancel := context.WithCancel(ctx)
	s.cancel = cancel

	conn, resp, err := websocket.DefaultDialer.Dial(addr, http.Header{})
	log.Println("connected to server:", resp)
	if err != nil {
		return err
	}
	s.socket = conn

	// run heartbeat and listen routines
	go s.heartbeat(ctx, time.Duration(s.heartbeatInterval*uint(time.Second)))
	go s.listen(ctx)

	return nil
}

// disconnect closes and cleans up the connection.
func (s *socket) disconnect() error {
	defer s.cancel()
	return s.socket.Close()
}

// push sends data on the connection.
func (s *socket) push(data interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.socket.WriteJSON(data)
}

// heartbeat is a routine informing the connection the client is still alive.
func (s *socket) heartbeat(ctx context.Context, interval time.Duration) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg := Message{
				Topic: PHOENIX_TOPIC,
				Event: EVENT_HEARTBEAT,
			}
			err := s.push(msg)
			if err != nil {
				log.Println("heartbeat error:", err)
			}
			time.Sleep(interval)
		}
	}
}

// listen is a routine receiving messages from the connection.
func (s *socket) listen(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			var message Message
			if err := s.socket.ReadJSON(&message); err != nil {
				log.Println("message read error:", err)
			}

			// handle events and route messages
			switch message.Event {
			case EVENT_REPLY:
			case EVENT_JOIN:
			case EVENT_MESSAGE:
			default:
				s.router.RouteMessage(&message)
			}
		}
	}
}
