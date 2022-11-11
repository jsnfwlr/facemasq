package events

import (
	"regexp"

	"facemasq/lib/events/safe"
)

// Filter compiles the string as a regex and returns the original listener (lstnr) wrapped in a new listener that filters incoming events by the Kind
func Filter(match string, original Listener) Listener {
	if match == "" || match == "*" {
		return original
	}
	regex := regexp.MustCompile(match)
	return func(event Event) {
		if regex.MatchString(event.Kind) {
			safe.Run(func() {
				original(event)
			})
		}
	}
}
