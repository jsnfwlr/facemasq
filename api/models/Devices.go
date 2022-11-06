package models

import (
	"time"

	"github.com/volatiletech/null"
)

type Devices []Device

type Device struct {
	ID                int64       `bun:",notnull,pk,autoincrement" json:"ID"`
	MachineName       string      `bun:",type:varchar(64),nullzero,notnull,default:'Unknown'" json:"MachineName"`
	Brand             null.String `bun:",type:varchar(64)" json:"Brand"`
	Model             null.String `bun:",type:varchar(64)" json:"Model"`
	Purchased         null.String `bun:",type:varchar(10)" json:"Purchased"`
	Serial            null.String `bun:",type:varchar(128)" json:"Serial"`
	FirstSeen         time.Time   `bun:",type:timestamp,nullzero,notnull" json:"FirstSeen"`
	IsTracked         null.Bool   `bun:",type:boolean,default:false" json:"IsTracked"`
	IsGuest           null.Bool   `bun:",type:boolean,default:false" json:"IsGuest"`
	IsOnline          null.Bool   `bun:",type:boolean,default:false" json:"IsOnline"`
	Label             null.String `bun:",type:varchar(64),notnull" json:"Label"`
	Notes             null.String `bun:",type:text" json:"Notes"`
	CategoryID        int64       `bun:",nullzero,notnull,default:1" json:"CategoryID"`
	StatusID          int64       `bun:",nullzero,notnull,default:1" json:"StatusID"`
	MaintainerID      int64       `bun:",nullzero,notnull,default:1" json:"MaintainerID"`
	LocationID        int64       `bun:",nullzero,notnull,default:1" json:"LocationID"`
	DeviceTypeID      int64       `bun:",nullzero,notnull,default:1" json:"DeviceTypeID"`
	OperatingSystemID int64       `bun:",nullzero,notnull,default:1" json:"OperatingSystemID"`
	ArchitectureID    int64       `bun:",nullzero,notnull,default:1" json:"ArchitectureID"`
	// CreatedAt         time.Time         `bun:",nullzero,notnull,default:current_timestamp"`
	// DeletedAt         time.Time         `bun:",soft_delete,nullzero"`
	Interfaces []Interface       `bun:"-" json:"Interfaces"`
	SortOrder  string            `bun:"-"`
	Primary    PrimaryConnection `bun:"-"`
}

func GetDeviceTestData() (seed []Device) {
	seed = []Device{
		{
			Label:     null.String{String: "Testing 1", Valid: true},
			Notes:     null.String{String: "Testing", Valid: true},
			FirstSeen: time.Now(),
		},
		{
			Label:     null.String{String: "Testing 2", Valid: true},
			Notes:     null.String{String: "Testing", Valid: true},
			FirstSeen: time.Now(),
		},
		{
			Label:     null.String{String: "Testing 3", Valid: true},
			Notes:     null.String{String: "Testing", Valid: true},
			FirstSeen: time.Now(),
		},
	}
	return
}
