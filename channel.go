package realtimego

import (
	"fmt"
	"log"
)

// ChannelOption represents the channel configuration options.
// i.e. WithTable
type ChannelOption func(ch *Channel)

// Channel manages a subscription to a realtime socket.
type Channel struct {
	client *Client

	topic Topic

	// OnInsert is a message handler for INSERT event messages
	OnInsert func(Message)
	// OnUpdate is a message handler for UPDATE event messages
	OnUpdate func(Message)
	// OnDelete is a message handler for DELETE event messages
	OnDelete func(Message)
}

// newChannel returns a channel used to subscribe and unsubscribe to topics.
func newChannel(c *Client, options ...ChannelOption) (*Channel, error) {
	ch := &Channel{
		client: c,
	}

	// set default message handlers as needed
	ch.setDefaultMessageHandlers()

	// set options
	for _, opt := range options {
		opt(ch)
	}

	if ch.topic == "" {
		return ch, fmt.Errorf("cannot create a channel with empty topic:%v", ch.topic)
	}

	return ch, nil
}

// Subscribe requests to receive messages for a topic from the realtime server.
func (ch *Channel) Subscribe() error {
	// add to router
	ch.client.router.AddChannel(ch)

	log.Println("subscribing to topic:", ch.topic)
	msg := Message{
		Topic:   ch.topic,
		Event:   EVENT_JOIN,
		Payload: ch.client.params,
	}

	return ch.client.socket.push(msg)
}

// Unsubscribe requests to stop receiving messages for a topic from the realtime server.
func (ch *Channel) Unsubscribe() error {
	// remove from router
	ch.client.router.DelChannel(ch)

	log.Println("unsubscribing from topic:", ch.topic)
	msg := Message{
		Topic:   ch.topic,
		Event:   EVENT_LEAVE,
		Payload: ch.client.params,
	}

	return ch.client.socket.push(msg)
}

// WithTable option sets the database/schema/table for a channel.
// passing a nil for a following parameter will construct a topic without that target.
// i.e. ("mydatabase", "myschema", nil) == "mydatabase:myschema"
// the first nil encountered stops the construction of the topic.
// i.e. (nil, "myschema", "mytable") == ""
func WithTable(database *string, schema *string, table *string) ChannelOption {
	return func(ch *Channel) {
		topic := ""

		if database == nil {
			ch.topic = Topic(topic)
			return
		}
		topic += *database

		if schema == nil {
			ch.topic = Topic(topic)
			return
		}
		topic += fmt.Sprintf(":%s", *schema)

		if table == nil {
			ch.topic = Topic(topic)
			return
		}
		topic += fmt.Sprintf(":%s", *table)

		ch.topic = Topic(topic)
	}

}

func (ch *Channel) setDefaultMessageHandlers() {
	if ch.OnInsert == nil {
		ch.OnInsert = func(m Message) {
			fmt.Println("INSERT:", m)
		}
	}

	if ch.OnUpdate == nil {
		ch.OnUpdate = func(m Message) {
			fmt.Println("UPDATE:", m)
		}
	}

	if ch.OnDelete == nil {
		ch.OnDelete = func(m Message) {
			fmt.Println("DELETE:", m)
		}
	}
}
