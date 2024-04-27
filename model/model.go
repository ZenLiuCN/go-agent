package model

import (
	"database/sql"
	"github.com/ZenLiuCN/gofra/hasher"
	"github.com/ZenLiuCN/gofra/modeler"
)

// go:generate go install github.com/ZenLiuCN/gofra/gofra-gene

//go:generate gofra-gene -a -t User
type User struct {
	modeler.FullModelEntity[int64]
	Nick   string         `db:"nick" json:"nick"`
	Name   string         `db:"name" json:"name"`
	Secret sql.NullString `db:"secret" json:"secret"`
	Totp   sql.NullString `db:"totp" json:"totp"`
}

func (e *User) CheckTOTP(code string) bool {
	if !e.Totp.Valid {
		return false
	}
	return hasher.TotpValidate(code, e.Totp.String)
}
func (e *User) CheckSecret(raw string) bool {
	if !e.Secret.Valid {
		return false
	}
	return hasher.PasswordValidate(raw, e.Secret.String)
}
