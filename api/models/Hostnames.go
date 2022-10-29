package models

import (
	"github.com/volatiletech/null"
)

type Hostnames []Hostname

type Hostname struct {
	ID        int64       `bun:",notnull,pk,autoincrement"`
	Hostname  string      `bun:",type:varchar(256),notnull,unique"`
	IsDNS     bool        `bun:",type:boolean,nullzero,notnull,default:false"`
	IsSelfSet bool        `bun:",type:boolean,nullzero,notnull,default:false"`
	Notes     null.String `bun:",type:text"`
	AddressID int64       `bun:""`
}

func GetHostnameTestData(addresses []Address) (seed []Hostname) {
	seed = []Hostname{
		{
			Hostname:  "Host0",
			AddressID: addresses[0].ID,
			IsDNS:     true,
			IsSelfSet: false,
		},
		{
			Hostname:  "Host1",
			AddressID: addresses[1].ID,
			IsDNS:     true,
			IsSelfSet: false,
		},
	}
	return
}
