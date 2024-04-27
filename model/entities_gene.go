package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ZenLiuCN/gofra/modeler"
	"time"
)

const (

	// UserNick generate for User.Nick <nick::string>
	UserNick modeler.FIELD = 9

	// UserName generate for User.Name <name::string>
	UserName modeler.FIELD = 10

	// UserSecret generate for User.Secret <secret::sql.NullString>
	UserSecret modeler.FIELD = 11

	// UserTotp generate for User.Totp <totp::sql.NullString>
	UserTotp modeler.FIELD = 12
)

var (
	UserFields modeler.EntityInfo[int64, User]
)

func init() {

	UserFields = modeler.EntityInfoBuilder[int64, User](

		func(u *User, f modeler.FIELD) any {
			switch f {
			case modeler.FIELD_ID:
				return u.Id
			case modeler.FIELD_CREATE_AT:
				return u.CreateAt
			case modeler.FIELD_MODIFIED_AT:
				return u.ModifiedAt
			case modeler.FIELD_REMOVED:
				return u.Removed
			case modeler.FIELD_VERSION:
				return u.Version
			case modeler.FIELD_CREATE_BY:
				return u.CreateBy
			case modeler.FIELD_MODIFIED_BY:
				return u.ModifiedBy

			case UserNick:
				return u.Nick

			case UserName:
				return u.Name

			case UserSecret:
				return u.Secret

			case UserTotp:
				return u.Totp

			default:
				panic(fmt.Errorf("invalid field %d", f))
			}
		},
		func(u *User, f modeler.FIELD, v any) {
			switch f {
			case modeler.FIELD_ID:
				if x, ok := v.(int64); ok {
					u.Id = x
				} else {
					panic(fmt.Errorf("bad field %T type of %d", v, f))
				}
			case modeler.FIELD_CREATE_AT:
				if x, ok := v.(time.Time); ok {
					u.CreateAt = x
				} else {
					panic(fmt.Errorf("bad field %T type of %d", v, f))
				}
			case modeler.FIELD_MODIFIED_AT:
				if x, ok := v.(time.Time); ok {
					u.ModifiedAt = x
				} else {
					panic(fmt.Errorf("bad field %T type of %d", v, f))
				}
			case modeler.FIELD_REMOVED:
				if x, ok := v.(bool); ok {
					u.Removed = x
				} else {
					panic(fmt.Errorf("bad field %T type of %d", v, f))
				}
			case modeler.FIELD_VERSION:
				if x, ok := v.(int); ok {
					u.Version = x
				} else {
					panic(fmt.Errorf("bad field %T type of %d", v, f))
				}
			case modeler.FIELD_CREATE_BY:
				if x, ok := v.(int64); ok {
					u.CreateBy = x
				} else {
					panic(fmt.Errorf("bad field %T type of %d", v, f))
				}
			case modeler.FIELD_MODIFIED_BY:
				if x, ok := v.(int64); ok {
					u.ModifiedBy = x
				} else {
					panic(fmt.Errorf("bad field %T type of %d", v, f))
				}

			case UserNick:

				if x, ok := v.(string); ok {
					u.Nick = x
				} else {
					panic(fmt.Errorf("bad field %T type of %d", v, f))
				}
			case UserName:

				if x, ok := v.(string); ok {
					u.Name = x
				} else {
					panic(fmt.Errorf("bad field %T type of %d", v, f))
				}
			case UserSecret:

				if x, ok := v.(sql.NullString); ok {
					u.Secret = x
				} else {
					panic(fmt.Errorf("bad field %T type of %d", v, f))
				}
			case UserTotp:

				if x, ok := v.(sql.NullString); ok {
					u.Totp = x
				} else {
					panic(fmt.Errorf("bad field %T type of %d", v, f))
				}
			default:
				panic(fmt.Errorf("invalid field %d", f))
			}
		},
		map[modeler.FIELD]string{
			modeler.FIELD_ID:          "id",
			modeler.FIELD_CREATE_AT:   "create_at",
			modeler.FIELD_MODIFIED_AT: "modified_at",
			modeler.FIELD_REMOVED:     "removed",
			modeler.FIELD_VERSION:     "version",
			modeler.FIELD_CREATE_BY:   "create_by",
			modeler.FIELD_MODIFIED_BY: "modified_by",
			UserNick:                  "nick",
			UserName:                  "name",
			UserSecret:                "secret",
			UserTotp:                  "totp",
		},
	)
}

type UserEntity struct {
	User
	modeler.BaseEntity[int64, User]
}

func newUserEntity(tab string, configurer modeler.Configurer, s User, executor modeler.Executor) (e *UserEntity) {
	e = &UserEntity{
		User: s,
	}
	e.BaseEntity = modeler.NewBaseEntity[int64, User](
		configurer,
		&e.User,
		UserFields,
		executor,
		modeler.NewBaseSQLMaker[int64, User](tab, configurer, UserFields),
		func(ctx context.Context) bool {
			return true
		},
	)
	return
}

func (c *UserEntity) GetNick() string {
	return c.DoRead(UserNick).(string)
}
func (c *UserEntity) SetNick(v string) bool {
	return c.DoWrite(UserNick, v)
}

func (c *UserEntity) GetName() string {
	return c.DoRead(UserName).(string)
}
func (c *UserEntity) SetName(v string) bool {
	return c.DoWrite(UserName, v)
}

func (c *UserEntity) GetSecret() sql.NullString {
	return c.DoRead(UserSecret).(sql.NullString)
}
func (c *UserEntity) SetSecret(v sql.NullString) bool {
	return c.DoWrite(UserSecret, v)
}

func (c *UserEntity) GetTotp() sql.NullString {
	return c.DoRead(UserTotp).(sql.NullString)
}
func (c *UserEntity) SetTotp(v sql.NullString) bool {
	return c.DoWrite(UserTotp, v)
}
