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
	AccessLevel     int64       `bun:",notnull,default:1" json:"-"`
	CanAuthenticate bool        `bun:",type:boolean,nullzero,notnull,default:false" json:"-"`
	IsInternal      bool        `bun:",type:boolean,nullzero,notnull,default:false"`
	IsLocked        bool        `bun:",type:boolean,nullzero,notnull,default:false"`
	// CreatedAt       time.Time   `bun:",nullzero,notnull,default:current_timestamp"`
	// DeletedAt       time.Time   `bun:",soft_delete,nullzero"`
}
