package models

import (
	"github.com/volatiletech/null"
)

type Locations []Location

type Location struct {
	ID       int64       `bun:",notnull,pk,autoincrement"`
	Label    string      `bun:",type:varchar(64),unique,notnull"`
	Notes    null.String `bun:",type:text"`
	IsCloud  bool        `bun:",type:boolean,notnull,default:false"`
	IsLocked bool        `bun:",type:boolean,notnull,default:false"`
}

func GetLocationSeed() (seed []Location) {
	seed = []Location{
		{
			Label:    "Limbo",
			IsLocked: true,
		},
	}
	return
}
