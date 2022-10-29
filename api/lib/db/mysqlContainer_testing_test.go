//go:build full || testing

package db

import (
	"testing"
	"time"
)

func TestMySQLContainer(t *testing.T) {
	mysql, err := StartMySQLContainer("test", "mysqlUser", "pass", "3306")
	if err != nil {
		t.Fatalf("unable to start mysql container %v", err)
	}

	t.Cleanup(func() {
		mysql.Cleanup()
	})

	t.Log("Sleeping for 30 seconds")
	time.Sleep(30 * time.Second)
	t.Log("Cleaning up")

	err = mysql.Cleanup()
	if err != nil {
		t.Fatalf("unable to clean mysql container %v", err)
	}

}
