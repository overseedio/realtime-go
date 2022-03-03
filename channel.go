package realtimego

// Channel manages a subscription to a realtime socket.
type Channel struct{}

// Subscribe requests to receive messages for a topic from the realtime server.
func (c *Channel) Subscribe() error {
	return nil
}

// Unsubscribe requests to stop receiving messages for a topic from the realtime server.
func (c *Channel) Unsubscribe() error {
	return nil
}
