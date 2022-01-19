package History

import "github.com/jsnfwlr/facemasq/api/lib/db"

// import (
// 	"database/sql"

// 	"github.com/jsnfwlr/facemasq/api/lib/db"
// )

const TABLENAME = `History`

type Models []Model

type Model struct {
	AddressID int `db:"AddressID"`
	ScanID    int `db:"ScanID"`
}

func Get() (records []Model, err error) {
	sql := `SELECT * FROM ` + TABLENAME + `;`
	err = db.Conn.Select(&records, sql)
	return
}
