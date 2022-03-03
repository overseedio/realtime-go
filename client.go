package realtimego

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// Client manages connections to the realtime server topics.
type Client struct {
	addr, apiKey string

	socket *socket

	heartbeatInterval uint
}

// NewClient returns a realtime client.
func NewClient(addr, apiKey string, options ...func(*Client)) (*Client, error) {
	// convert address to ws
	addr, err := addressToWebsocket(addr)
	if err != nil {
		return nil, err
	}

	// add api key param if provided
	if apiKey != "" {
		addr = fmt.Sprintf("%v?apikey=%v&vsn=1.0.0", addr, apiKey)
	}

	c := &Client{
		addr:   addr,
		apiKey: apiKey,
	}

	// set options
	for _, opt := range options {
		opt(c)
	}

	// create socket
	socket := newSocket(c.heartbeatInterval)
	c.socket = socket

	return c, nil
}

// Connect creates a connection to the server.
func (c *Client) Connect() error {
	return c.socket.connect(context.Background(), c.addr)
}

// Disconnect closes and cleans up the client connection.
func (c *Client) Disconnect() error {
	return c.socket.disconnect()
}

// Channel creates a new subscription channel to the realtime server.
func (c *Client) Channel() *Channel {
	return nil
}

// SetAuth updates the client and channels with the latest token.
func (c *Client) SetAuth(token string) {}

func addressToWebsocket(addr string) (string, error) {
	_, err := url.ParseRequestURI(addr)
	if err != nil {
		return "", err
	}

	// replace protocol
	addr = strings.ReplaceAll(addr, "http", "ws")

	addr = fmt.Sprintf("%v/realtime/v1/websocket", addr)

	return addr, nil
}

func WithHeartbeatInterval(interval uint) func(*Client) {
	return func(c *Client) {
		c.heartbeatInterval = interval
	}
}
