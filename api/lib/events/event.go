package events

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (e Event) String() string {
	b, _ := e.MarshalJSON()

	return string(b)
}

// MarshalJSON implements the json marshaler for an event
func (e Event) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"kind": e.Kind,
	}
	if len(e.Message) != 0 {
		m["message"] = e.Message
	}
	if e.Error != nil {
		m["error"] = e.Error.Error()
	}
	if len(e.Payload) != 0 {
		m["payload"] = e.Payload
	}

	return json.Marshal(m)
}

// Validate that an event is ready to be emitted
func (e Event) Validate() error {
	if len(e.Kind) == 0 {
		return fmt.Errorf("kind can not be blank")
	}
	return nil
}

func (e Event) IsError() bool {
	return strings.HasSuffix(e.Kind, ":err")
}
