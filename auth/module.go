package auth

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/gofra/hasher"
)

var (
	//go:embed secret.d.ts
	secretDefine  []byte
	secretDeclare = map[string]any{
		"password": hasher.PasswordValidate,
		"totp":     hasher.TotpValidate,
		"bcrypt":   hasher.BcryptHash,
		"argon2":   hasher.Argon2Hash,
		"totpGen":  hasher.TotpGenerate,
		"totpCode": hasher.TotpCode,
	}
	secret = secretModule{}
)

func init() {
	engine.RegisterModule(secret)

}

type secretModule struct {
}

func (s secretModule) TypeDefine() []byte {
	return secretDefine
}

func (s secretModule) Identity() string {
	return "agent/secret"
}

func (s secretModule) Exports() map[string]any {
	return secretDeclare
}
