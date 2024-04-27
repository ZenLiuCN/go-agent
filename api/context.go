package api

import (
	"github.com/ZenLiuCN/go-agent/model"
	"github.com/ZenLiuCN/gofra/utils"
)

type Request struct {
	User *model.User
	utils.Cache[string, any]
}
