package realtimego

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// ClientOption represents the client configuration options.
// i.e. WithTable
type ClientOption func(c *Client)

// Client manages connections to the realtime server topics.
type Client struct {
	addr, apiKey string
	params       map[string]interface{}

	socket *socket
	router *router

	heartbeatInterval uint
}

// NewClient returns a realtime client.
func NewClient(addr, apiKey string, options ...ClientOption) (*Client, error) {
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
		params: map[string]interface{}{},
	}

	// set options
	for _, opt := range options {
		opt(c)
	}

	// create router
	router := newRouter()
	c.router = router

	// create socket
	socket := newSocket(c.heartbeatInterval)
	c.socket = socket
	c.socket.router = router

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
func (c *Client) Channel(options ...ChannelOption) (*Channel, error) {
	return newChannel(c, options...)
}

// SetAuth updates the client and channels with the latest token.
func (c *Client) SetAuth(token string) {
	c.params[PARAM_USER_TOKEN] = token
}

// WithHeartbeatInterval option sets the heartbeat interval () on the socket connection.
func WithHeartbeatInterval(interval uint) func(*Client) {
	return func(c *Client) {
		c.heartbeatInterval = interval
	}
}

// WithUserToken option sets the user_token parameter for user auth when communicating with the server.
// i.e. authenticating the user for an RLS protected table.
func WithUserToken(token string) ClientOption {
	return func(c *Client) {
		c.params[PARAM_USER_TOKEN] = token
	}
}

// WithParams option sets the request parameters used when sending data to the server.
func WithParams(params map[string]interface{}) ClientOption {
	// avoid nil
	if params == nil {
		params = map[string]interface{}{}
	}

	return func(c *Client) {
		c.params = params
	}
}

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
