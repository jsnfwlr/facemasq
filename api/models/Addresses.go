package models

import (
	"time"

	"github.com/volatiletech/null"
)

type Addresses []Address

type Address struct {
	ID           int64         `bun:",notnull,pk,autoincrement" json:"ID"`
	IPv4         null.String   `bun:"ipv4,type:varchar(15),unique:UC_Address" json:"IPv4"`
	IPv6         null.String   `bun:"ipv6,type:varchar(64)" json:"IPv6"`
	IsPrimary    null.Bool     `bun:",default:true" json:"IsPrimary"`
	IsVirtual    null.Bool     `bun:",default:false" json:"IsVirtual"`
	IsReserved   null.Bool     `bun:",default:false" json:"IsReserved"`
	LastSeen     time.Time     `bun:",nullzero,notnull,default:current_timestamp" json:"LastSeen"`
	Label        null.String   `bun:",type:varchar(64)" json:"Label"`
	Notes        null.String   `bun:",type:text" json:"Notes"`
	InterfaceID  int64         `bun:",notnull,unique:UC_Address" json:"InterfaceID"`
	Interface    *Interface    `bun:"rel:belongs-to,join:interface_id=id"`
	Hostnames    []Hostname    `bun:"-" json:"Hostnames"`
	Connectivity []Connections `bun:"-" json:"Connectivity"`
	SortOrder    string        `bun:"-"`
}

type Connectionss []Connections

type Connections struct {
	State bool      `bun:"state"`
	Time  time.Time `bun:"time"`
}

type ConnectionGroups []ConnectionGroup

type ConnectionGroup struct {
	Time        time.Time `bun:"time"`
	AddressList string    `bun:"addresses"`
}

func (addresses Addresses) GetInterfaceIDs() (ids []int64) {
	for i := range addresses {
		ids = append(ids, addresses[i].InterfaceID)
	}
	return
}
