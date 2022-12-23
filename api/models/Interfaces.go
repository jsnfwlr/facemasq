package models

import (
	"time"

	"github.com/volatiletech/null"
)

type Interfaces []Interface

type Interface struct {
	ID        int64       `bun:",notnull,pk,autoincrement" json:"ID"`
	MAC       string      `bun:",type:varchar(17),notnull,unique" json:"MAC"`
	IsPrimary bool        `bun:",type:boolean,notnull,default:true" json:"IsPrimary"`
	IsVirtual bool        `bun:",type:boolean,notnull,default:false" json:"IsVirtual"`
	IsOnline  bool        `bun:",type:boolean,notnull,default:false" json:"IsOnline"`
	Label     null.String `bun:",type:varchar(64)" json:"Label"`
	Notes     null.String `bun:",type:text" json:"Notes"`
	LastSeen  time.Time   `bun:",nullzero,notnull,default:current_timestamp" json:"LastSeen"`
	// CreatedAt       time.Time         `bun:",nullzero,notnull,default:current_timestamp"`
	// DeletedAt       time.Time         `bun:",soft_delete,nullzero"`
	StatusID        int64             `bun:",nullzero,notnull,default:1" json:"StatusID"`
	InterfaceTypeID int64             `bun:",default:1,notnull" json:"InterfaceTypeID"`
	VlanID          int64             `bun:",nullzero,notnull,default:1" json:"VLANID"`
	DeviceID        int64             `bun:",notnull" json:"DeviceID"`
	Device          *Device           `bun:"rel:belongs-to,join:device_id=id"`
	Addresses       []Address         `bun:"-" json:"Addresses"`
	SortOrder       string            `bun:"-"`
	Primary         PrimaryConnection `bun:"-"`
}

func GetInterfaceTestData(devices []Device, lastSeen time.Time) (seed []Interface) {
	seed = []Interface{
		{
			MAC:             "00:00:00:00:00:00",
			IsPrimary:       true,
			IsVirtual:       false,
			InterfaceTypeID: 1,
			LastSeen:        lastSeen,
			VlanID:          1,
			DeviceID:        devices[0].ID,
			Label:           null.String{String: "Test", Valid: true},
			Notes:           null.String{String: "Testing", Valid: true},
		},
		{
			MAC:             "00:00:00:00:00:01",
			IsPrimary:       true,
			IsVirtual:       false,
			InterfaceTypeID: 1,
			LastSeen:        lastSeen,
			VlanID:          1,
			DeviceID:        devices[1].ID,
			Label:           null.String{String: "Test", Valid: true},
			Notes:           null.String{String: "Testing", Valid: true},
		},
		{
			MAC:             "00:00:00:00:01:00",
			IsPrimary:       true,
			IsVirtual:       false,
			InterfaceTypeID: 1,
			LastSeen:        lastSeen,
			VlanID:          1,
			DeviceID:        devices[2].ID,
			Label:           null.String{String: "Test", Valid: true},
			Notes:           null.String{String: "Testing", Valid: true},
		},
		{
			MAC:             "00:00:00:00:01:01",
			IsPrimary:       true,
			IsVirtual:       false,
			InterfaceTypeID: 1,
			LastSeen:        lastSeen,
			VlanID:          1,
			DeviceID:        devices[2].ID,
			Label:           null.String{String: "Test", Valid: true},
			Notes:           null.String{String: "Testing", Valid: true},
		},
	}
	return
}
