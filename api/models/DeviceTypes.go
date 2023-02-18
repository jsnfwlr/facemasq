package models

import (
	"github.com/volatiletech/null"
)

type DeviceTypes []DeviceType

type DeviceType struct {
	ID       int64       `bun:",notnull,pk,autoincrement"`
	Label    string      `bun:",type:varchar(64),unique,notnull"`
	Icon     string      `bun:",type:varchar(64),notnull"`
	Notes    null.String `bun:",type:text"`
	IsLocked bool        `bun:",type:boolean,notnull,default:false"`
}

func GetDeviceTypeSeed() (seed []DeviceType) {
	seed = []DeviceType{
		{
			Label:    "Unspecified",
			Icon:     "HelpCircle",
			IsLocked: true,
		},
	}
	return
}

// var _ bun.BeforeAppendModelHook = (*DeviceType)(nil)

// func (m *DeviceType) BeforeAppendModel(ctx context.Context, query bun.Query) (err error) {
// 	// switch query.(type) {
// 	// case *bun.InsertQuery:
// 	// 	m.CreatedAt = time.Now()
// 	// case *bun.UpdateQuery:
// 	// 	m.UpdatedAt = time.Now()
// 	// }

// 	return
// }
