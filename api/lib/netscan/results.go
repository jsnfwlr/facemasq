package netscan

import "github.com/volatiletech/null"

type Result struct {
	ScanID    int64
	Hostname  string
	IPv4      string
	IPv6      null.String
	MAC       string
	FirstSeen string
	LastSeen  string
	ScanCount int
	Notes     string
}
