package Scans

import "github.com/jsnfwlr/facemasq/api/lib/db"

// import (
// 	"database/sql"

// 	"github.com/jsnfwlr/facemasq/api/lib/db"
// )

const TABLENAME = `Scans`

type Models []Model

type Model struct {
	ID   int64  `db:"ID"`
	Time string `db:"Time"`
}

func Get() (records []Model, err error) {
	sql := `SELECT * FROM ` + TABLENAME + `;`
	err = db.Conn.Select(&records, sql)
	return
}

// func (records Models) Save() {
// 	for i := range records {
// 		records[i].Save()
// 	}
// }

// func (record *Model) Save() (err error) {
// 	var result sql.Result
// 	if record.ID == 0 {
// 		sql := `INSERT INTO Scans (Label, Notes) VALUES (?,?);`
// 		result, err = db.Conn.Exec(sql, record.Label, record.Notes)
// 		if err != nil {
// 			return
// 		}
// 		record.ID, err = result.LastInsertId()
// 	} else {
// 		sql := `UPDATE ` + TABLENAME +` SET Label = ?, Notes = ? WHERE ID = ?;`
// 		_, err = db.Conn.Exec(sql, record.Label, record.Notes, record.ID)
// 	}
// 	return
// }

// func (record *Model) Delete() (err error) {
// 	sql := `DELETE FROM ` + TABLENAME +` WHERE ID = ?;`
// 	_, err = db.Conn.Exec(sql, record.ID)
// 	return
// }
