package api

import (
	"github.com/ZenLiuCN/go-agent/model"
)

type Request struct {
	User *model.User
	units.Cache[string, any]
}
