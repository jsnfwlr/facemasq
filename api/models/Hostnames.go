package models

import (
	"github.com/volatiletech/null"
)

type Hostnames []Hostname

type Hostname struct {
	ID        int64       `bun:",notnull,pk,autoincrement"`
	Hostname  string      `bun:",type:varchar(256),notnull,unique"`
	IsDNS     bool        `bun:",nullzero,notnull,default:false"`
	IsSelfSet bool        `bun:",nullzero,notnull,default:false"`
	Notes     null.String `bun:",type:text"`
	AddressID int64       `bun:""`
}
