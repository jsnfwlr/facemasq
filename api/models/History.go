package models

type Histories []History

type History struct {
	AddressID int `bun:",notnull,unique:UC_History"`
	ScanID    int `bun:",notnull,unique:UC_History"`
}
