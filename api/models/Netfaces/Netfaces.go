package Netfaces

import (
	"database/sql"
	"time"

	"github.com/volatiletech/null"

	"github.com/jsnfwlr/facemasq/api/lib/db"
	"github.com/jsnfwlr/facemasq/api/models/Addresses"
	"github.com/jsnfwlr/facemasq/api/models/PrimaryConnection"
)

const TABLENAME = `Interfaces`

type Models []Model

type Model struct {
	ID              int64             `db:"ID" json:"ID"`
	MAC             string            `db:"MAC" json:"MAC"`
	IsPrimary       bool              `db:"IsPrimary" json:"IsPrimary"`
	IsVirtual       bool              `db:"IsVirtual" json:"IsVirtual"`
	IsOnline        bool              `db:"IsOnline" json:"IsOnline"`
	LastSeen        time.Time         `db:"LastSeen" json:"LastSeen"`
	Label           null.String       `db:"Label" json:"Label"`
	Notes           null.String       `db:"Notes" json:"Notes"`
	StatusID        int64             `db:"StatusID" json:"StatusID"`
	InterfaceTypeID int64             `db:"InterfaceTypeID" json:"InterfaceTypeID"`
	VLANID          int64             `db:"VLANID" json:"VLANID"`
	DeviceID        int64             `db:"DeviceID" json:"DeviceID"`
	Addresses       []Addresses.Model `json:"Addresses"`
	SortOrder       string
	Primary         PrimaryConnection.Model
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
		sql := `INSERT INTO ` + TABLENAME + ` (MAC, IsPrimary, IsVirtual, IsOnline, LastSeen, Label, Notes, InterfaceTypeID, VLANID, StatusID, DeviceID) VALUES (?,?,?,?,?,?,?,?,?,?,?);`
		result, err = db.Conn.Exec(sql, record.MAC, record.IsPrimary, record.IsVirtual, record.IsOnline, record.LastSeen, record.Label, record.Notes, record.InterfaceTypeID, record.VLANID, record.StatusID, record.DeviceID)
		if err != nil {
			return
		}
		record.ID, err = result.LastInsertId()
	} else {
		sql := `UPDATE ` + TABLENAME + ` SET MAC = ?, IsPrimary = ?, IsVirtual = ?, IsOnline = ?, LastSeen = ?, Label = ?, Notes = ?, InterfaceTypeID = ?, VLANID = ?, DeviceID = ?, StatusID = ? WHERE ID = ? AND MAC = ?;`
		_, err = db.Conn.Exec(sql, record.MAC, record.IsPrimary, record.IsVirtual, record.IsOnline, record.LastSeen, record.Label, record.Notes, record.InterfaceTypeID, record.VLANID, record.DeviceID, record.StatusID, record.ID, record.MAC)
	}
	return
}

func (record *Model) Delete() (err error) {
	sql := `DELETE FROM ` + TABLENAME + ` WHERE ID = ?;`
	_, err = db.Conn.Exec(sql, record.ID)
	return
}
