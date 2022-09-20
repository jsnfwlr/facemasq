package models

import (
	"github.com/volatiletech/null"
)

type OperatingSystems []OperatingSystem

type OperatingSystem struct {
	ID           int64       `bun:",notnull,pk,autoincrement"`
	Vendor       string      `bun:",type:varchar(64),notnull,unique:US_OperatingSystem"`
	Family       string      `bun:",type:varchar(64),notnull,unique:US_OperatingSystem"`
	Version      string      `bun:",type:varchar(64),notnull,unique:US_OperatingSystem"`
	Name         string      `bun:",type:varchar(64),notnull,unique:US_OperatingSystem"`
	IsOpenSource bool        `bun:",nullzero,notnull,default:false"`
	IsServer     bool        `bun:",nullzero,notnull,default:false"`
	Notes        null.String `bun:",type:text"`
	IsLocked     bool        `bun:",nullzero,notnull,default:false"`
}

func GetOperatingSystemSeed() (seed []OperatingSystem) {
	seed = []OperatingSystem{
		{
			Vendor:       "?",
			Family:       "?",
			Version:      "?",
			Name:         "?",
			IsOpenSource: false,
			IsServer:     false,
			IsLocked:     true,
		},
		{
			Vendor:       "Apple",
			Family:       "MacOS",
			Version:      "10.13",
			Name:         "High Sierra",
			IsOpenSource: false,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Apple",
			Family:       "MacOS",
			Version:      "10.14",
			Name:         "Mojave",
			IsOpenSource: false,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Apple",
			Family:       "MacOS",
			Version:      "10.15",
			Name:         "Catalina",
			IsOpenSource: false,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Apple",
			Family:       "MacOS",
			Version:      "11",
			Name:         "Big Sur",
			IsOpenSource: false,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Apple",
			Family:       "MacOS",
			Version:      "12",
			Name:         "Monterey",
			IsOpenSource: false,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Microsoft",
			Family:       "Windows",
			Version:      "10",
			Name:         "Win10",
			IsOpenSource: false,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Microsoft",
			Family:       "Windows",
			Version:      "11",
			Name:         "Win11",
			IsOpenSource: false,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Canonical",
			Family:       "Ubuntu",
			Version:      "18.04",
			Name:         "Bionic Beaver",
			IsOpenSource: true,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Canonical",
			Family:       "Ubuntu",
			Version:      "20.04",
			Name:         "Focal Fossa",
			IsOpenSource: true,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Canonical",
			Family:       "Ubuntu",
			Version:      "21.04",
			Name:         "Hirsute Hippo",
			IsOpenSource: true,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Canonical",
			Family:       "Ubuntu",
			Version:      "21.10",
			Name:         "Impish Indri",
			IsOpenSource: true,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Canonical",
			Family:       "Ubuntu",
			Version:      "22.04",
			Name:         "Jammy Jellyfish",
			IsOpenSource: true,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Google",
			Family:       "Android",
			Version:      "6",
			Name:         "Marshmallow",
			IsOpenSource: true,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Google",
			Family:       "Android",
			Version:      "7",
			Name:         "Nougat",
			IsOpenSource: true,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Google",
			Family:       "Android",
			Version:      "8",
			Name:         "Oreo",
			IsOpenSource: true,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Google",
			Family:       "Android",
			Version:      "9",
			Name:         "Pie",
			IsOpenSource: true,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Google",
			Family:       "Android",
			Version:      "10",
			Name:         "Q",
			IsOpenSource: true,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Google",
			Family:       "Android",
			Version:      "11",
			Name:         "R",
			IsOpenSource: true,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Google",
			Family:       "Android",
			Version:      "12",
			Name:         "S",
			IsOpenSource: true,
			IsServer:     false,
			IsLocked:     false,
		},
		{
			Vendor:       "Google",
			Family:       "Android",
			Version:      "13",
			Name:         "T",
			IsOpenSource: true,
			IsServer:     false,
			IsLocked:     false,
		},
	}
	return
}
