package Hostnames

import (
	"database/sql"

	"github.com/volatiletech/null"

	"github.com/jsnfwlr/facemasq/api/lib/db"
)

const TABLENAME = `Hostnames`

type Models []Model

type Model struct {
	ID        int64       `db:"ID"`
	Hostname  string      `db:"Hostname"`
	IsDNS     bool        `db:"IsDNS"`
	IsSelfSet bool        `db:"IsSelfSet"`
	Notes     null.String `db:"Notes"`
	AddressID int64       `db:"AddressID"`
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
		sql := `INSERT INTO Hostnames (Hostname, IsDNS, IsSelfSet, Notes, AddressID) VALUES (?,?,?,?,?);`
		result, err = db.Conn.Exec(sql, record.Hostname, record.IsDNS, record.IsSelfSet, record.Notes, record.AddressID)
		if err != nil {
			return
		}
		record.ID, err = result.LastInsertId()
	} else {
		sql := `UPDATE ` + TABLENAME + ` SET Hostname = ?, IsDNS = ?, IsSelfSet = ?, Notes = ?, AddressID = ? WHERE ID = ?;`
		_, err = db.Conn.Exec(sql, record.Hostname, record.IsDNS, record.IsSelfSet, record.Notes, record.AddressID, record.ID)
	}
	return
}

func (record *Model) Delete() (err error) {
	sql := `DELETE FROM ` + TABLENAME + ` WHERE ID = ?;`
	_, err = db.Conn.Exec(sql, record.ID)
	return
}
