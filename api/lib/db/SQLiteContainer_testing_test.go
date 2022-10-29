//go:build full || testing

package db

import (
	"testing"
	"time"
)

func TestSQLiteContainer(t *testing.T) {
	sqlite, err := StartSQLiteContainer("network.sqlite")
	if err != nil {
		t.Fatalf("unable to start sqlite dummy-container %v", err)
	}

	t.Cleanup(func() {
		sqlite.Cleanup()
	})

	t.Log("Sleeping for 30 seconds")
	time.Sleep(30 * time.Second)
	t.Log("Cleaning up")

	err = sqlite.Cleanup()
	if err != nil {
		t.Fatalf("unable to clean sqlite dummy-container %v", err)
	}
}
