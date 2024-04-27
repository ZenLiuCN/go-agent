package units

import (
	"github.com/ZenLiuCN/gofra/utils"
	"time"
)

type Cached[K comparable, V any] struct {
	utils.Cache[K, V]
}

func (c Cached[K, V]) Close() {
	c.Cache.StopKeeping()
}
func NewStringCache(freq, ttl time.Duration, units utils.MeasureUnit, opt ...utils.Option) (c *Cached[string, any]) {
	c = &Cached[string, any]{
		Cache: utils.NewCache[string, any]("", nil, freq, ttl, units, opt...),
	}
	c.StartKeeping()
	return
}
func NewNumberCache(freq, ttl time.Duration, units utils.MeasureUnit, opt ...utils.Option) (c *Cached[float64, any]) {
	c = &Cached[float64, any]{
		Cache: utils.NewCache[float64, any](0, nil, freq, ttl, units, opt...),
	}
	c.StartKeeping()
	return
}
