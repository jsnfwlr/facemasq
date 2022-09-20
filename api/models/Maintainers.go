package models

import (
	"github.com/uptrace/bun"
	"github.com/volatiletech/null"
)

type Maintainers []Maintainer

type Maintainer struct {
	bun.BaseModel   `bun:"table:users,alias:maintainers"`
	ID              int64       `bun:",notnull,pk,autoincrement"`
	Username        null.String `bun:"" json:"-"`
	Password        null.String `bun:"" json:"-"`
	Label           string      `bun:",unique,notnull"`
	Notes           null.String `bun:",type:text"`
	CanAuthenticate bool        `bun:"" json:"-"`
	AccessLevel     int64       `bun:",notnull,default:1" json:"-"`
	IsInternal      bool        `bun:",notnull,default:false"`
	IsLocked        bool        `bun:",nullzero,notnull,default:false"`
}
