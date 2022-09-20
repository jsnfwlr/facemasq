package models

import (
	"time"

	"github.com/volatiletech/null"
)

type Interfaces []Interface

type Interface struct {
	ID              int64             `bun:",notnull,pk,autoincrement" json:"ID"`
	MAC             string            `bun:",type:varchar(17),notnull,unique" json:"MAC"`
	IsPrimary       bool              `bun:",nullzero,notnull,default:true" json:"IsPrimary"`
	IsVirtual       bool              `bun:",nullzero,notnull,default:false" json:"IsVirtual"`
	IsOnline        bool              `bun:",nullzero,notnull,default:false" json:"IsOnline"`
	Label           null.String       `bun:",type:varchar(64)" json:"Label"`
	Notes           null.String       `bun:",type:text" json:"Notes"`
	LastSeen        time.Time         `bun:",nullzero,notnull,default:current_timestamp" json:"LastSeen"`
	StatusID        int64             `bun:",nullzero,notnull,default:1" json:"StatusID"`
	InterfaceTypeID int64             `bun:",default:1,notnull" json:"InterfaceTypeID"`
	VlanID          int64             `bun:",nullzero,notnull,default:1" json:"VLANID"`
	DeviceID        int64             `bun:",notnull" json:"DeviceID"`
	Addresses       []Address         `bun:"-" json:"Addresses"`
	SortOrder       string            `bun:"-"`
	Primary         PrimaryConnection `bun:"-"`
}
