//go:build test
// +build test

package db

import (
	"testing"
	"time"
)

func TestMySQL(t *testing.T) {
	if mysql, err := StartMySQLDB("test", "postgres", "pass", "3306"); err != nil {
		t.Fatalf("unable to start postgres container %v", err)
	} else {
		time.Sleep(30 * time.Second)
		if err := mysql.Cleanup(); err != nil {
			t.Fatalf("unable to clean postgres container %v", err)
		}
	}
}
