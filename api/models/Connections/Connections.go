package Connections

import "time"

type Model struct {
	State bool      `db:"State"`
	Time  time.Time `db:"Time"`
}
