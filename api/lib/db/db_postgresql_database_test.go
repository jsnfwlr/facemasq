//go:build full || (database && postgresql) || (database && !mysql && !postgresql && !sqlite)

package db

import "testing"

func init() {
	DBEngines = append(DBEngines, "postgres")
}

func TestPostgreSQL(t *testing.T) {
	DBEngine = "postgres"
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
