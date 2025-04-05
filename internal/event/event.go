package event

type Event int

const (
	EventConnectToServer Event = iota
	EventStartServer
)
