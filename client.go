package realtimego

// Client manages connections to the realtime server topics.
type Client struct{}

// NewClient returns a realtime client.
func NewClient() *Client {
	return nil
}

// Connect creates a connection to the server.
func (c *Client) Connect() error {
	return nil
}

// Disconnect closes and cleans up the client connection.
func (c *Client) Disconnect() error {
	return nil
}

// Channel creates a new subscription channel to the realtime server.
func (c *Client) Channel() *Channel {
	return nil
}

// SetAuth updates the client and channels with the latest token.
func (c *Client) SetAuth(token string) {}
