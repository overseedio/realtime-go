package realtimego

type Event string

const (
	EVENT_MESSAGE        Event = "phx_message"
	EVENT_JOIN           Event = "phx_join"
	EVENT_LEAVE          Event = "phx_leave"
	EVENT_REPLY          Event = "phx_reply"
	EVENT_HEARTBEAT      Event = "heartbeat"
	EVENT_MESSAGE_INSERT Event = "INSERT"
	EVENT_MESSAGE_UPDATE Event = "UPDATE"
	EVENT_MESSAGE_DELETE Event = "DELETE"
)

type Topic string

const (
	PHOENIX_TOPIC Topic = "phoenix"
)
