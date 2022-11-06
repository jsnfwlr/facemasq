package main

import "testing"

func TestMessage(t *testing.T) {
	err := SendMessage("New Device", "Blah", 1)
	if err != nil {
		t.Error(err)
	}
}
