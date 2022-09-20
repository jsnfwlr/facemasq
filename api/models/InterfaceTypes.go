package models

import (
	"github.com/volatiletech/null"
)

type InterfaceTypes []InterfaceType

type InterfaceType struct {
	ID       int64       `bun:",notnull,pk,autoincrement"`
	Label    string      `bun:",type:varchar(64),unique,notnull"`
	Icon     string      `bun:",type:varchar(64),notnull"`
	Notes    null.String `bun:",type:text"`
	IsLocked bool        `bun:",nullzero,notnull,default:false"`
}

func GetInterfaceTypeSeed() (seed []InterfaceType) {
	seed = []InterfaceType{
		{
			Label:    "WiFi",
			Icon:     "HelpCircle",
			IsLocked: true,
		},
		{
			Label:    "Ethernet Cable",
			Icon:     "HelpCircle",
			IsLocked: true,
		},

		{
			Label:    "Fibre",
			Icon:     "HelpCircle",
			IsLocked: true,
		},

		{
			Label:    "Internal",
			Icon:     "HelpCircle",
			IsLocked: true,
		},
	}
	return
}
