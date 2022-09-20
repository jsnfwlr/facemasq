package models

import (
	"github.com/volatiletech/null"
)

type Metas []Meta

type Meta struct {
	Name   string     `bun:",type:varchar(256),notnull,unique:UC_Meta" json:"Name"`
	Value  string     `bun:",type:text,notnull" json:"Value"`
	UserID null.Int64 `bun:",unique:UC_Meta" json:"UserID"`
}
