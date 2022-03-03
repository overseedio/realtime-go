package realtimego

type Event string

const (
	EVENT_MESSAGE   Event = "phx_message"
	EVENT_JOIN      Event = "phx_join"
	EVENT_LEAVE     Event = "phx_leave"
	EVENT_HEARTBEAT Event = "heartbeat"
)

type Topic string

const (
	PHOENIX_TOPIC Topic = "phoenix"
)
