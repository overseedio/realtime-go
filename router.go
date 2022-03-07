package realtimego

import (
	"sync"
)

type router struct {
	channels map[Topic]*Channel
	mu       sync.RWMutex
}

func newRouter() *router {
	return &router{
		channels: make(map[Topic]*Channel),
	}
}

func (r *router) AddChannel(ch *Channel) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.channels[ch.topic] = ch
}

func (r *router) DelChannel(ch *Channel) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.channels, ch.topic)
}

func (r *router) RouteMessage(msg *Message) {
	r.mu.RLock()

	// get topic to route message to channel
	ch, ok := r.channels[msg.Topic]
	if !ok {
		r.mu.RUnlock()
		return
	}
	r.mu.RUnlock()

	// route message to channel handler
	switch msg.Event {
	case EVENT_MESSAGE_INSERT:
		ch.OnInsert(*msg)
	case EVENT_MESSAGE_UPDATE:
		ch.OnUpdate(*msg)
	case EVENT_MESSAGE_DELETE:
		ch.OnDelete(*msg)
	}

}
