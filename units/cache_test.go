package units

import (
	"github.com/ZenLiuCN/engine"
	"github.com/ZenLiuCN/fn"
	"testing"
)

func TestCache(t *testing.T) {
	v := engine.Get()
	defer v.Free()
	fn.Panic1(v.RunTs(
		//language=typescript
		`
import {newNumberKeyCache, newStringKeyCache, withMaxSize} from 'agent/units'
import * as time from 'go/time'
const cache=newStringKeyCache(time.duration('10s'),time.duration('15s'),2,withMaxSize(10))
registerResource(cache)
cache.put("123",{a:1,b:1})
const x=cache.get("123")
console.log(x[0],x[1])

const cache2=newNumberKeyCache(time.duration('10s'),time.duration('15s'),2,withMaxSize(10))
registerResource(cache)
cache2.put(123,{a:1,b:1})
const x2=cache2.get(123)
console.log(x2[0],x2[1])
console.log(time.toDuration(2,time.Second).string())
console.log(time.fromDuration(time.Minute,time.Second))
`))
}
