package models

import (
	"github.com/volatiletech/null"
)

type Architectures []Architecture

type Architecture struct {
	ID       int64       `bun:",notnull,pk,autoincrement"`
	Label    string      `bun:",type:varchar(64),unique,notnull"`
	BitSpace int64       `bun:",nullzero,notnull,default:64"`
	Notes    null.String `bun:",type:text"`
	IsLocked bool        `bun:",type:boolean,nullzero,notnull,default:false"`
}

func GetArchitectureSeed() (seed []Architecture) {
	seed = []Architecture{
		{
			Label:    "Unknown",
			BitSpace: 0,
			IsLocked: true,
		},
		{
			Label:    "x86",
			BitSpace: 32,
			IsLocked: false,
		},
		{
			Label:    "x64",
			BitSpace: 64,
			IsLocked: false,
		},
		{
			Label:    "ARM",
			BitSpace: 32,
			IsLocked: false,
		},
		{
			Label:    "ARM64",
			BitSpace: 64,
			IsLocked: false,
		},
		{
			Label:    "RISC-V 32I",
			BitSpace: 32,
			IsLocked: false,
		},
		{
			Label:    "RISC-V 32E",
			BitSpace: 32,
			IsLocked: false,
		},
		{
			Label:    "RISC-V 64I",
			BitSpace: 64,
			IsLocked: false,
		},
		{
			Label:    "RISC-V 128I",
			BitSpace: 128,
			IsLocked: false,
		},
	}
	return
}
