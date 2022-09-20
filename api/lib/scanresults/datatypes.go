package scanresults

import (
	"github.com/volatiletech/null"
)

type Record struct {
	ScanID      int64       `bun:"-"`
	Hostname    string      `bun:"-"`
	IPv4        string      `bun:"IPv4"`
	IPv6        null.String `bun:"-"`
	MAC         string      `bun:"MAC"`
	FirstSeen   string      `bun:"-"`
	LastSeen    string      `bun:"-"`
	ScanCount   int         `bun:"-"`
	Notes       string      `bun:"-"`
	DeviceID    int64       `bun:"-"`
	InterfaceID int64       `bun:"-"`
	AddressID   int64       `bun:"-"`
}

type Records []Record
