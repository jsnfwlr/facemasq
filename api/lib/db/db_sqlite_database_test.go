//go:build full || (database && sqlite) || (database && !mysql && !postgresql && !sqlite)

package db

import (
	"testing"
)

func init() {
	DBEngines = append(DBEngines, "sqlite")
}

func TestSQLite(t *testing.T) {
	DBEngine = "sqlite"
	cntnr, err := ConnectToTest(true)
	if err != nil {
		t.Fatal(err)
	}
	RunningContainer = cntnr

	t.Cleanup(func() {
		t.Log("Removing container")
		RunningContainer.Close()
	})

}
