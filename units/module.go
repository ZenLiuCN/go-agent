package units

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/gofra/units"
)

var (
	//go:embed units.d.ts
	define  []byte
	declare = map[string]any{
		"withMaxSize":           units.WithMaximize,
		"withExpireAfterAccess": units.WithExpiredAfterAccess,
		"newStringKeyCache":     NewStringCache,
		"newNumberKeyCache":     NewNumberCache,
	}
	model = module{}
)

func init() {
	engine.RegisterModule(model)
}

type module struct {
}

func (s module) TypeDefine() []byte {
	return define
}

func (s module) Identity() string {
	return "agent/units"
}

func (s module) Exports() map[string]any {
	return declare
}
