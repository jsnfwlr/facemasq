package events

import "facemasq/lib/events/mapi"

type Payload = mapi.Mapi

// Event represents different events in the lifecycle of a Buffalo app
type Event struct {
	// Kind is the "type" of event "app:start"
	Kind string `json:"kind"`
	// Message is optional
	Message string `json:"message"`
	// Payload is optional
	Payload Payload `json:"payload"`
	// Error is optional
	Error error `json:"-"`
}

// Listener is a function capable of handling events
type Listener func(e Event)

type DeleteFn func()

// Manager can be implemented to replace the default events manager
type Manager interface {
	Listen(string, Listener) (DeleteFn, error)
	Emit(Event) error
}

type manager struct {
	listeners listenerMap
}
