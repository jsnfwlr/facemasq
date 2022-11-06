package extensions

import "facemasq/lib/events"

type Listener struct {
	Kind     string
	Listener events.Listener
}
