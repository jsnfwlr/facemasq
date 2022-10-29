package models

import (
	"github.com/volatiletech/null"
)

type DeviceTypes []DeviceType

type DeviceType struct {
	ID       int64       `bun:",notnull,pk,autoincrement"`
	Label    string      `bun:",type:varchar(64),unique,notnull"`
	Icon     string      `bun:",type:varchar(64),notnull"`
	Notes    null.String `bun:",type:text"`
	IsLocked bool        `bun:",type:boolean,nullzero,notnull,default:false"`
}

func GetDeviceTypeSeed() (seed []DeviceType) {
	seed = []DeviceType{
		{
			Label:    "Unspecified",
			Icon:     "HelpCircle",
			IsLocked: true,
		},
	}
	return
}
