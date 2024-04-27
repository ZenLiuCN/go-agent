package api

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"
)

var (
	//go:embed api.d.ts
	apiDefine  []byte
	apiDeclare = map[string]any{}
	api        = apiModule{}
)

func init() {
	engine.RegisterModule(api)
}

type apiModule struct {
}

func (a apiModule) TypeDefine() []byte {
	return apiDefine
}

func (a apiModule) Identity() string {
	return "agent/api"
}

func (a apiModule) Exports() map[string]any {
	return apiDeclare
}
