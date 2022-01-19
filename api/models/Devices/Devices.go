package Devices

import (
	"database/sql"

	"github.com/volatiletech/null"

	"github.com/jsnfwlr/facemasq/api/lib/db"
	"github.com/jsnfwlr/facemasq/api/models/Netfaces"
	"github.com/jsnfwlr/facemasq/api/models/PrimaryConnection"
)

const TABLENAME = `Devices`

type Models []Model

type Model struct {
	ID                int64            `db:"ID" json:"ID"`
	MachineName       string           `db:"MachineName" json:"MachineName"`
	Brand             null.String      `db:"Brand" json:"Brand"`
	Model             null.String      `db:"Model" json:"Model"`
	Purchased         null.String      `db:"Purchased" json:"Purchased"`
	Serial            null.String      `db:"Serial" json:"Serial"`
	IsTracked         bool             `db:"IsTracked" json:"IsTracked"`
	FirstSeen         null.String      `db:"FirstSeen" json:"FirstSeen"`
	IsGuest           bool             `db:"IsGuest" json:"IsGuest"`
	IsOnline          bool             `db:"IsOnline" json:"IsOnline"`
	Label             null.String      `db:"Label" json:"Label"`
	Notes             null.String      `db:"Notes" json:"Notes"`
	CategoryID        int64            `db:"CategoryID" json:"CategoryID"`
	StatusID          int64            `db:"StatusID" json:"StatusID"`
	MaintainerID      int64            `db:"MaintainerID" json:"MaintainerID"`
	LocationID        int64            `db:"LocationID" json:"LocationID"`
	DeviceTypeID      int64            `db:"DeviceTypeID" json:"DeviceTypeID"`
	OperatingSystemID int64            `db:"OperatingSystemID" json:"OperatingSystemID"`
	ArchitectureID    int64            `db:"ArchitectureID" json:"ArchitectureID"`
	Interfaces        []Netfaces.Model `json:"Interfaces"`
	SortOrder         string
	Primary           PrimaryConnection.Model
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
		sql := `INSERT INTO ` + TABLENAME + ` (MachineName, Brand,        Model,        Purchased,        Serial,        IsTracked,        FirstSeen,        IsGuest,        IsOnline       , Label,        Notes,        CategoryID,        StatusID,        MaintainerID,        LocationID,        DeviceTypeID,        OperatingSystemID,        ArchitectureID) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
		result, err = db.Conn.Exec(sql, record.MachineName, record.Brand, record.Model, record.Purchased, record.Serial, record.IsTracked, record.FirstSeen, record.IsGuest, record.IsOnline, record.Label, record.Notes, record.CategoryID, record.StatusID, record.MaintainerID, record.LocationID, record.DeviceTypeID, record.OperatingSystemID, record.ArchitectureID)
		if err != nil {
			return
		}
		record.ID, err = result.LastInsertId()
	} else {
		sql := `UPDATE ` + TABLENAME + ` SET
		                           MachineName = ?, 	 Brand = ?,    Model = ?,    Purchased = ?,    Serial = ?,    IsTracked = ?,    FirstSeen = ?,    IsGuest = ?,    IsOnline = ?,    Label = ?,    Notes = ?,    CategoryID = ?,    StatusID = ?,    MaintainerID = ?,    LocationID = ?,    DeviceTypeID = ?,    OperatingSystemID = ?,    ArchitectureID = ?     WHERE ID = ?;`
		_, err = db.Conn.Exec(sql, record.MachineName, record.Brand, record.Model, record.Purchased, record.Serial, record.IsTracked, record.FirstSeen, record.IsGuest, record.IsOnline, record.Label, record.Notes, record.CategoryID, record.StatusID, record.MaintainerID, record.LocationID, record.DeviceTypeID, record.OperatingSystemID, record.ArchitectureID, record.ID)
	}
	return
}

func (record *Model) Delete() (err error) {
	sql := `DELETE FROM ` + TABLENAME + ` WHERE ID = ?;`
	_, err = db.Conn.Exec(sql, record.ID)
	return
}
