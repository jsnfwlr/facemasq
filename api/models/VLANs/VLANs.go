package VLANs

import (
	"database/sql"

	"github.com/volatiletech/null"

	"github.com/jsnfwlr/facemasq/api/lib/db"
)

const TABLENAME = `VLANs`

type Models []Model

type Model struct {
	ID       int64       `db:"ID"`
	Label    string      `db:"Label"`
	IPv4Mask string      `db:"IPv4Mask"`
	IPv6Mask string      `db:"IPv6Mask"`
	Notes    null.String `db:"Notes"`
	IsLocked bool        `db:"IsLocked"`
}

func Get() (records []Model, err error) {
	sql := `SELECT * FROM ` + TABLENAME + `;`
	err = db.Conn.Select(&records, sql)
	return
}

func (records Models) Save() {
	for i := range records {
		records[i].Save()
	}
}

func (record *Model) Save() (err error) {
	var result sql.Result
	if record.ID == 0 {
		sql := `INSERT INTO VLANs (Label, IPv4Mask, IPv6Mask, Notes) VALUES (?,?,?,?);`
		result, err = db.Conn.Exec(sql, record.Label, record.IPv4Mask, record.IPv6Mask, record.Notes)
		if err != nil {
			return
		}
		record.ID, err = result.LastInsertId()
	} else {
		sql := `UPDATE ` + TABLENAME + ` SET Label = ?, IPv4Mask = ?, IPv6Mask = ?, Notes = ? WHERE ID = ?;`
		_, err = db.Conn.Exec(sql, record.Label, record.IPv4Mask, record.IPv6Mask, record.Notes, record.ID)
	}
	return
}

func (record *Model) Delete() (err error) {
	var match Model
	sql := `SELECT * FROM VLANs WHERE ID = ? AND IsLocked = 0;`
	err = db.Conn.Get(&match, sql, record.ID)
	if err != nil {
		return
	}
	sql = `DELETE FROM ` + TABLENAME + ` WHERE ID = ?;`
	_, err = db.Conn.Exec(sql, record.ID)
	return
}
