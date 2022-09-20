package models

type Historys []History

type History struct {
	AddressID int `bun:",notnull,unique:UC_History"`
	ScanID    int `bun:",notnull,unique:UC_History"`
}
