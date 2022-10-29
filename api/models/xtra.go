package models

import "time"

type (
	Connections        []Connection
	ConnectionGroups   []ConnectionGroup
	PrimaryConnections []PrimaryConnection
)

type Connection struct {
	State bool      `bun:"state"`
	Time  time.Time `bun:"time"`
}

type ConnectionGroup struct {
	Time        time.Time `bun:"time"`
	AddressList string    `bun:"addresses"`
}

type PrimaryConnection struct {
	IPv4            string
	IPv6            string
	MAC             string
	VlanID          int64
	InterfaceTypeID int64
	IsReservedIP    bool
	IsVirtualIP     bool
	IsVirtualIFace  bool
}
