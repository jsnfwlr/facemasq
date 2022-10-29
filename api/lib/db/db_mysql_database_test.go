//go:build full || (database && mysql) || (database && !mysql && !postgresql && !sqlite)

package db

import "testing"

func init() {
	DBEngines = append(DBEngines, "mysql")
}

func TestMySQL(t *testing.T) {
	DBEngine = "mysql"
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
