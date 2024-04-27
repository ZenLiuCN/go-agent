package model

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"
)

var (
	//go:embed model.d.ts
	modelDefine  []byte
	modelDeclare = map[string]any{
		"newUserStore": NewUserStore,
	}
	model = modelModule{}
)

func init() {
	engine.RegisterModule(model)
}

type modelModule struct {
}

func (s modelModule) TypeDefine() []byte {
	return modelDefine
}

func (s modelModule) Identity() string {
	return "agent/model"
}

func (s modelModule) Exports() map[string]any {
	return modelDeclare
}
