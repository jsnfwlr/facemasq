package models

import (
	"time"

	"github.com/volatiletech/null"
)

// Addresses are rows of the addresses table
type Addresses []Address

// Address is a row of the addresses table
type Address struct {
	ID           int64        `bun:",notnull,pk,autoincrement" json:"ID"`
	IPv4         null.String  `bun:"ipv4,type:varchar(15),unique:UC_Address" json:"IPv4"`
	IPv6         null.String  `bun:"ipv6,type:varchar(64)" json:"IPv6"`
	IsPrimary    null.Bool    `bun:",type:boolean,default:true" json:"IsPrimary"`
	IsVirtual    null.Bool    `bun:",type:boolean,default:false" json:"IsVirtual"`
	IsReserved   null.Bool    `bun:",type:boolean,default:false" json:"IsReserved"`
	LastSeen     time.Time    `bun:",nullzero,notnull,default:current_timestamp" json:"LastSeen"`
	Label        null.String  `bun:",type:varchar(64)" json:"Label"`
	Notes        null.String  `bun:",type:text" json:"Notes"`
	InterfaceID  int64        `bun:",notnull,unique:UC_Address" json:"InterfaceID"`
	Interface    *Interface   `bun:"rel:belongs-to,join:interface_id=id"`
	Hostnames    []Hostname   `bun:"-" json:"Hostnames"`
	Connectivity []Connection `bun:"-" json:"Connectivity"`
	SortOrder    string       `bun:"-"`
}

func (addresses Addresses) GetInterfaceIDs() (ids []int64) {
	for i := range addresses {
		ids = append(ids, addresses[i].InterfaceID)
	}
	return
}

func GetAddressTestData(netfaces []Interface, lastSeen time.Time) (seed []Address) {
	seed = []Address{
		{
			IPv4:        null.String{String: "192.168.0.1", Valid: true},
			IPv6:        null.String{String: ""},
			IsPrimary:   null.BoolFrom(true),
			IsVirtual:   null.BoolFrom(false),
			IsReserved:  null.BoolFrom(false),
			LastSeen:    lastSeen,
			Label:       null.String{String: "Test", Valid: true},
			Notes:       null.String{String: "This should be deleted", Valid: true},
			InterfaceID: netfaces[0].ID,
		},
		{
			IPv4:        null.String{String: "192.168.0.2", Valid: true},
			IPv6:        null.String{String: ""},
			IsPrimary:   null.BoolFrom(true),
			IsVirtual:   null.BoolFrom(false),
			IsReserved:  null.BoolFrom(false),
			LastSeen:    lastSeen,
			Label:       null.String{String: "Test", Valid: true},
			Notes:       null.String{String: "This should be deleted", Valid: true},
			InterfaceID: netfaces[1].ID,
		},
		{
			IPv4:        null.String{String: "192.168.1.1", Valid: true},
			IPv6:        null.String{String: ""},
			IsPrimary:   null.BoolFrom(true),
			IsVirtual:   null.BoolFrom(false),
			IsReserved:  null.BoolFrom(false),
			LastSeen:    lastSeen,
			Label:       null.String{String: "Test", Valid: true},
			Notes:       null.String{String: "This should be deleted", Valid: true},
			InterfaceID: netfaces[2].ID,
		},
		{
			IPv4:        null.String{String: "192.168.1.2", Valid: true},
			IPv6:        null.String{String: ""},
			IsPrimary:   null.BoolFrom(true),
			IsVirtual:   null.BoolFrom(false),
			IsReserved:  null.BoolFrom(false),
			LastSeen:    lastSeen,
			Label:       null.String{String: "Test", Valid: true},
			Notes:       null.String{String: "This should be deleted", Valid: true},
			InterfaceID: netfaces[3].ID,
		},
		{
			IPv4:        null.String{String: "192.168.1.3", Valid: true},
			IPv6:        null.String{String: ""},
			IsPrimary:   null.BoolFrom(true),
			IsVirtual:   null.BoolFrom(false),
			IsReserved:  null.BoolFrom(false),
			LastSeen:    lastSeen,
			Label:       null.String{String: "Test", Valid: true},
			Notes:       null.String{String: "This should be deleted", Valid: true},
			InterfaceID: netfaces[3].ID,
		},
	}
	return
}
