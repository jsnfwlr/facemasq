package db

import (
	"os"

	"github.com/jmoiron/sqlx"
)

const DBTargetVer = 1

var adminPassword string
var Conn *sqlx.DB

type Engine struct {
    EngineType string
    DataRoot string
    DataFile string
    HostName string
    Username string
    Password string
}

/*
type Engine interface {
    Connect() error
    prepare() error
    checkVersion() error
    upgrade(newVersion int) error
}
*/

func (input Engine) Connect() (conn *sqlx.DB, err error) {
    switch input.EngineType {
	case "sqlite":
		engine := SQLite {
			DataRoot: input.DataRoot,
			DataFile: input.DataFile,
		}
		conn, err = engine.Connect()
//	case "mysql", "mariadb":
//	case "postgresql":
    }
    return
}

func init() {
	adminPassword = os.Getenv("ADMINPASSWORD")
	if adminPassword == "" {
		adminPassword = "ResetMe"
	}
}
