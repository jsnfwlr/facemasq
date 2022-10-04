//go:build (database && test) || full

package db

import (
	"testing"
	"time"
)

func TestPostgres(t *testing.T) {
	if pg, err := StartPostgresDB("test", "postgresUser", "pass", "5432"); err != nil {
		t.Fatalf("unable to start postgres container %v", err)
	} else {
		time.Sleep(30 * time.Second)
		if err := pg.Cleanup(); err != nil {
			t.Fatalf("unable to clean postgres container %v", err)
		}
	}
}
