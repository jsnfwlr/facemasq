//go:build full || testing

package db

import (
	"testing"
	"time"
)

func TestPostgreSQLContainer(t *testing.T) {
	pg, err := StartPostgreSQLContainer("test", "postgresUser", "pass", "5432")
	if err != nil {
		t.Fatalf("unable to start postgres container %v", err)
	}

	t.Cleanup(func() {
		pg.Cleanup()
	})

	t.Log("Sleeping for 30 seconds")
	time.Sleep(30 * time.Second)
	t.Log("Cleaning up")

	err = pg.Cleanup()
	if err != nil {
		t.Fatalf("unable to clean postgres container %v", err)
	}
}
