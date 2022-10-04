//go:build (database && test) || full

package db

import (
	"testing"
	"time"
)

func TestMySQL(t *testing.T) {
	if mysql, err := StartMySQLDB("test", "mysqlUser", "pass", "3306"); err != nil {
		t.Fatalf("unable to start mysql container %v", err)
	} else {
		time.Sleep(30 * time.Second)
		if err := mysql.Cleanup(); err != nil {
			t.Fatalf("unable to clean mysql container %v", err)
		}
	}
}
