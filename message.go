package realtimego

type Message struct {
	Topic   string      `json:"topic"`
	Event   Event       `json:"event"`
	Payload interface{} `json:"payload"`
	Ref     int64       `json:"ref"`
}
