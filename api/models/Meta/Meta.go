package Meta

import (
	"github.com/volatiletech/null"

	"github.com/jsnfwlr/facemasq/api/lib/db"
)

const TABLENAME = `Meta`

type Models []Model

type Model struct {
	Name   string     `db:"Name" json:"Name"`
	Value  string     `db:"Value" json:"Value"`
	UserID null.Int64 `db:"UserID" json:"UserID"`
}

func Get(userID null.Int64) (records []Model, err error) {
	sql := `SELECT * FROM ` + TABLENAME + ` WHERE UserID = ?;`
	err = db.Conn.Select(&records, sql, userID)
	return
}

func (records Models) Save() {
	for i := range records {
		records[i].Save()
	}
}

func (record *Model) Save() (err error) {
	var match Model
	sql := `SELECT * FROM ` + TABLENAME + ` WHERE UserID = ? AND Name = ?;`
	err = db.Conn.Get(&match, sql, record.UserID, record.Name)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			sql = `INSERT INTO ` + TABLENAME + ` (Name, Value, UserID) VALUES (?,?, ?);`
			_, err = db.Conn.Exec(sql, record.Name, record.Value, record.UserID)
		}
		return
	}
	sql = `UPDATE ` + TABLENAME + ` SET Value = ? WHERE UserID = ? AND Name = ?;`
	_, err = db.Conn.Exec(sql, record.Value, record.UserID, record.Name)
	if err != nil {
		return
	}
	return
}

func (record *Model) Delete() (err error) {
	sql := `DELETE FROM ` + TABLENAME + ` WHERE UserID = ? AND Name = ?;`
	_, err = db.Conn.Exec(sql, record.UserID, record.Name)
	return
}
