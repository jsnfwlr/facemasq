package events

import (
	"strings"
)

const (
	// ErrGeneral is emitted for general errors
	ErrGeneral = "general:err"
	// ErrPanic is emitted when a panic is recovered
	ErrPanic = "panic:err"
)

// Emit an event to all listeners
func Emit(e Event) error {
	return boss.Emit(e)
}

func EmitPayload(kind string, payload interface{}) error {
	return EmitError(kind, nil, payload)
}

func EmitError(kind string, err error, payload interface{}) error {
	if err != nil && !strings.HasSuffix(kind, ":err") {
		kind += ":err"
	}
	var pl Payload
	pl, ok := payload.(Payload)
	if !ok {
		pl = Payload{
			"data": payload,
		}
	}
	e := Event{
		Kind:    kind,
		Payload: pl,
		Error:   err,
	}
	return Emit(e)
}
