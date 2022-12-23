package models

import (
	"github.com/volatiletech/null"
)

type Statuses []Status

type Status struct {
	ID       int64       `bun:",notnull,pk,autoincrement"`
	Label    string      `bun:",type:varchar(64),unique,notnull"`
	Icon     string      `bun:",type:varchar(64),notnull"`
	Notes    null.String `bun:",type:text"`
	IsLocked bool        `bun:",type:boolean,notnull,default:false"`
}

func GetStatusSeed() (seed []Status) {
	seed = []Status{
		{
			Label:    "Invading",
			Icon:     "HelpCircle",
			IsLocked: true,
		},
	}
	return
}
