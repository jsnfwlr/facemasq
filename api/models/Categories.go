package models

import (
	"github.com/volatiletech/null"
)

type Categories []Category

type Category struct {
	ID       int64       `bun:",notnull,pk,autoincrement"`
	Label    string      `bun:",type:varchar(64),unique,notnull"`
	Icon     string      `bun:",type:varchar(64),notnull"`
	Notes    null.String `bun:",type:text"`
	IsLocked bool        `bun:",type:boolean,nullzero,notnull,default:false"`
}

func GetCategorySeed() (seed []Category) {
	seed = []Category{
		{
			Label:    "Unsorted",
			Icon:     "HelpCircle",
			IsLocked: true,
		},
	}
	return
}
