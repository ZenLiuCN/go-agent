package model

import (
	"context"
	"fmt"
	jsql "github.com/ZenLiuCN/engine/sqlx"
	"github.com/ZenLiuCN/gofra/modeler"
)

type UserStore struct {
	table string
	modeler.SqlxExecutor
}

func (s *UserStore) fac(u User) *UserEntity {
	return newUserEntity(s.table, modeler.ConfigurerAll, u, s)
}
func (s *UserStore) Close(ctx context.Context) bool {
	return true
}

func (s *UserStore) ByName(name string) *UserEntity {
	u := User{}
	if err := s.DB.Get(&u, fmt.Sprintf("SELECT * FROM %s WHERE name=? AND removed=false", s.table), name); err != nil {
		return nil
	}
	return s.fac(u)
}
func (s *UserStore) ById(id int64) *UserEntity {
	u := User{}
	if err := s.DB.Get(&u, fmt.Sprintf("SELECT * FROM %s WHERE id=?  AND removed=false", s.table), id); err != nil {
		return nil
	}
	return s.fac(u)
}
func NewUserStore(table string, lx *jsql.SQLx) *UserStore {
	return &UserStore{table: table, SqlxExecutor: modeler.SqlxExecutor{DB: lx.DB}}
}
