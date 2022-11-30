package db

import (
	"fmt"
	"os"

	"facemasq/lib/files"
	"facemasq/lib/logging"
)

type SQLiteContainer struct {
	ID         string
	Connection ConnectionParams
	Cleanup    func() error
	dbType     string
}

func StartSQLiteContainer(filename string) (testContainer *SQLiteContainer, err error) {
	var dataPath string
	dataPath, err = files.GetDir("data")
	if err != nil {
		return
	}
	dbFile := fmt.Sprintf("%[2]s%[1]c%[3]s", os.PathSeparator, dataPath, filename)
	cleanup := func() error {
		if files.FileExists(dbFile) {
			return os.Remove(dbFile)
		}
		return nil
	}
	logging.Debug("Dummy container is running %s", dbFile)

	testContainer = &SQLiteContainer{
		ID: filename,
		Connection: ConnectionParams{
			DBFile: dbFile,
		},
		Cleanup: cleanup,
		dbType:  "sqlite",
	}
	return
}

func (sqlite *SQLiteContainer) Close() error {
	return sqlite.Cleanup()
}

func (sqlite *SQLiteContainer) GetConnection() ConnectionParams {
	return sqlite.Connection
}
