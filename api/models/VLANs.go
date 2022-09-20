package models

import (
	"github.com/volatiletech/null"
)

type VLANs []VLAN

type VLAN struct {
	ID       int64       `bun:",notnull,pk,autoincrement"`
	Label    string      `bun:",type:varchar(64),unique,notnull"`
	Maskv4   string      `bun:",type:varchar(19),notnull"`
	Maskv6   string      `bun:",type:varchar(128),notnull"`
	Notes    null.String `bun:",type:text"`
	IsLocked bool        `bun:",nullzero,notnull,default:false"`
}

func GetVLANSeed() (seed []VLAN) {
	seed = []VLAN{
		{
			Label:    "Default",
			Maskv4:   "0.0.0.0", // change to scan target value
			IsLocked: true,
		},
	}
	return
}
