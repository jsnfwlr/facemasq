package Addresses

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/volatiletech/null"

	"github.com/jsnfwlr/facemasq/api/lib/db"
	"github.com/jsnfwlr/facemasq/api/models/Connections"
	"github.com/jsnfwlr/facemasq/api/models/Hostnames"
)

const TABLENAME = `Addresses`

type Model struct {
	ID           int64               `db:"ID" json:"ID"`
	IPv4         null.String         `db:"IPv4" json:"IPv4"`
	IPv6         null.String         `db:"IPv6" json:"IPv6"`
	IsPrimary    bool                `db:"IsPrimary" json:"IsPrimary"`
	IsVirtual    bool                `db:"IsVirtual" json:"IsVirtual"`
	IsReserved   bool                `db:"IsReserved" json:"IsReserved"`
	LastSeen     time.Time           `db:"LastSeen" json:"LastSeen"`
	Label        null.String         `db:"Label" json:"Label"`
	Notes        null.String         `db:"Notes" json:"Notes"`
	InterfaceID  int64               `db:"InterfaceID" json:"InterfaceID"`
	Hostnames    []Hostnames.Model   `json:"Hostnames"`
	Connectivity []Connections.Model `json:"Connectivity"`
	SortOrder    string
}

type Models []Model

func Get() (records []Model, err error) {
	sql := `SELECT * FROM ` + TABLENAME + `;`
	err = db.Conn.Select(&records, sql)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (records Models) Save() {
	for i := range records {
		records[i].Save()
		// for j := range record[i].Hostnames {
		// 	record[i].Hostnames.Save
		// }
	}
}

func (record *Model) Save() (err error) {
	var result sql.Result
	if record.ID == 0 {
		sql := `INSERT INTO ` + TABLENAME + ` (IPv4, IPv6, IsPrimary, IsVirtual, IsReserved, LastSeen, Label, Notes, InterfaceID) VALUES (?,?,?,?,?,?,?,?,?);`
		result, err = db.Conn.Exec(sql, record.IPv4, record.IPv6, record.IsPrimary, record.IsVirtual, record.IsReserved, record.LastSeen, record.Label, record.Notes, record.InterfaceID)
		if err != nil {
			return
		}
		record.ID, err = result.LastInsertId()
	} else {
		sql := `UPDATE ` + TABLENAME + ` SET IPv4 = ?, IPv6 = ?, IsPrimary = ?, IsVirtual = ?, IsReserved = ?, LastSeen = ?, Label = ?, Notes = ?, InterfaceID = ? WHERE ID = ?;`
		_, err = db.Conn.Exec(sql, record.IPv4, record.IPv6, record.IsPrimary, record.IsVirtual, record.IsReserved, record.LastSeen, record.Label, record.Notes, record.InterfaceID, record.ID)
	}
	return
}

func (record *Model) Delete() (err error) {
	sql := `DELETE FROM ` + TABLENAME + ` WHERE ID = ?;`
	_, err = db.Conn.Exec(sql, record.ID)
	return
}
