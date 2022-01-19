package Users

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/volatiletech/null"

	"github.com/jsnfwlr/facemasq/api/lib/db"
	"github.com/jsnfwlr/facemasq/api/lib/password"
)

const TABLENAME = `Users`

type Models []Model

type Model struct {
	ID              int64       `db:"ID"`
	Username        null.String `db:"Username"`
	Password        null.String `db:"Password" json:"-"`
	Label           string      `db:"Label"`
	Notes           null.String `db:"Notes"`
	CanAuthenticate bool        `db:"CanAuthenticate"`
	AccessLevel     int         `db:"AccessLevel"`
	IsInternal      bool        `db:"IsInternal"`
	IsLocked        bool        `db:"IsLocked"`
	NewPassword     null.String `db:"NewPassword"`
}

func Get() (records []Model, err error) {
	sql := `SELECT *, NULL AS NewPassword FROM ` + TABLENAME + ` WHERE (Username IS NOT NULL OR IsLocked = 0);`
	err = db.Conn.Select(&records, sql)
	if err != nil {
		err = errors.New(sql)
	}
	return
}

func GetByID(id int64) (record Model, err error) {
	sql := `SELECT * FROM ` + TABLENAME + ` WHERE ID = ?;`
	err = db.Conn.Get(&record, sql, id)
	if err != nil {
		err = fmt.Errorf("%s %d", sql, id)
	}
	return
}

func (records Models) Save() {
	for i := range records {
		records[i].Save()
	}
}

func (record *Model) Save() (err error) {
	var hash string
	if record.NewPassword.Valid && len(record.NewPassword.String) > 0 {
		hash, err = password.HashPassword(record.NewPassword.String)
		if err != nil {
			return
		}
		record.Password = null.StringFrom(hash)
	}
	var result sql.Result
	if record.ID == 0 && ((record.NewPassword.Valid && len(record.NewPassword.String) > 0) || !record.CanAuthenticate) {
		sql := `INSERT INTO Users (Username, Password, Label, Notes, CanAuthenticate, AccessLevel, IsInternal) VALUES (?,?,?,?,?,?,?);`
		result, err = db.Conn.Exec(sql, record.Username, record.Password, record.Label, record.Notes, record.CanAuthenticate, record.AccessLevel, record.IsInternal)
		if err != nil {
			return
		}
		record.ID, err = result.LastInsertId()
		record.NewPassword = null.NewString("", false)
	} else if record.NewPassword.Valid && len(record.NewPassword.String) > 0 {
		sql := `UPDATE ` + TABLENAME + ` SET Username = ?, Password = ?, Label = ?, Notes = ?, CanAuthenticate = ?, AccessLevel = ?, IsInternal = ? WHERE ID = ?;`
		_, err = db.Conn.Exec(sql, record.Username, record.Password, record.Label, record.Notes, record.CanAuthenticate, record.AccessLevel, record.IsInternal, record.ID)
		record.NewPassword = null.NewString("", false)
	} else {
		sql := `UPDATE ` + TABLENAME + ` SET Username = ?, Label = ?, Notes = ?, CanAuthenticate = ?, AccessLevel = ?, IsInternal = ? WHERE ID = ?;`
		_, err = db.Conn.Exec(sql, record.Username, record.Label, record.Notes, record.CanAuthenticate, record.AccessLevel, record.IsInternal, record.ID)
	}
	return
}

func (record *Model) Delete() (err error) {
	var match Model
	match, err = GetByID(record.ID)
	if err != nil {
		return
	}
	if match.IsLocked {
		err = errors.New("record is locked - can not delete")
		return
	}
	sql := `DELETE FROM ` + TABLENAME + ` WHERE ID = ?;`
	_, err = db.Conn.Exec(sql, record.ID)
	return
}

func (record *Model) VerifyPassword(input string) (err error) {
	var newHash string
	if !record.Password.Valid {
		err = errors.New("stored password is invalid")
		return
	}
	newHash, err = password.ConfirmPassword(input, record.Password.String)
	if err != nil {
		return
	}
	if newHash != "" {
		record.NewPassword = null.StringFrom(newHash)
		err = record.Save()
	}
	return
}
