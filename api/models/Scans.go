package models

import (
	"time"
)

type Scans []Scan

type Scan struct {
	ID   int64     `bun:",notnull,pk,autoincrement"`
	Time time.Time `bun:",notnull,unique"`
}
