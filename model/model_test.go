package model

import (
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestModel(t *testing.T) {
	v := engine.Get()
	defer v.Free()
	u := new(User)
	v.Set("user", u)
	fn.Panic1(v.RunTs(
		//language=typescript
		`
		import type {User} from "agent/model"
		import * as sec from "agent/secret"
			const u:User=user
			console.log(u)
			console.log(u.totp)
			console.log(u.secret)
			console.log(u.secret.valid)
			console.log(u.secret.string)
			console.log(u.nick)
			console.log(u.name)
			console.log(u.id)     
			console.log(u.version)     
			console.log(u.removed)     
			console.log(u.createAt)     
			console.log(u.createBy)     
			console.log(u.modifiedAt)     
			console.log(u.modifiedBy)     
			u.secret.valid=true
			u.secret.string="$2a$10$.JTMHNxOzT4G9ybFfC/lVORewqqxCs/KDsrnWyTsvJZyIxFL/ZPSG"
			console.log("secret check",u.checkSecret("12345"))
			u.totp.string='otpauth://totp/some:user?algorithm=SHA512&digits=6&issuer=some&period=30&secret=EON4ELXYRBT7F3LCDHIXPR24MWGCJJMP'
			u.totp.valid=true
			console.log("totp check",u.checkTOTP(sec.totpCode(u.totp.string)))
		`))
}
