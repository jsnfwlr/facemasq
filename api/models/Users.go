package models

import (
	"facemasq/lib/logging"
	"facemasq/lib/password"

	"github.com/volatiletech/null"
)

type Users []User

type User struct {
	ID              int64       `bun:",notnull,pk,autoincrement"`
	Username        null.String `bun:",type:varchar(64),unique,nullzero"`
	Password        null.String `bun:",type:varchar(256)" json:"-"`
	Label           string      `bun:",type:varchar(64),unique,notnull"`
	CanAuthenticate bool        `bun:",nullzero,notnull,default:false"`
	AccessLevel     int64       `bun:",nullzero,notnull,default:0"`
	IsInternal      bool        `bun:",nullzero,notnull,default:false"`
	Notes           null.String `bun:",type:text"`
	IsLocked        bool        `bun:",nullzero,notnull,default:false"`
	NewPassword     null.String `bun:"-"`
}

func GetUserSeed(adminPassword string) (seed []User) {
	hashedPW, err := password.HashPassword(adminPassword)
	if err != nil {
		logging.Panic(err)
	}
	seed = []User{
		{
			Label:    "Invader",
			IsLocked: true,
		},
		{
			Username:        null.String{String: "Admin", Valid: true},
			Password:        null.String{String: hashedPW, Valid: true},
			Label:           "Admin",
			CanAuthenticate: true,
			AccessLevel:     1,
			IsLocked:        true,
		},
	}
	return
}
