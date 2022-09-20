package models

type Ports []Port

type Port struct {
	AddressID int64  `bun:",notnull,unique:UC_Port"`
	ScanID    int64  `bun:",notnull,unique:UC_Port"`
	Protocol  string `bun:",type:varchar(16),notnull,unique:UC_Port"`
	Port      int64  `bun:",notnull,unique:UC_Port"`
}
