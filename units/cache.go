package units

import (
	"time"
)

type Cached[K comparable, V any] struct {
	units.Cache[K, V]
}

func (c Cached[K, V]) Close() error {
	c.Cache.StopKeeping()
	return nil
}
func NewStringCache(freq, ttl time.Duration, units units.MeasureUnit, opt ...units.Option) (c *Cached[string, any]) {
	c = &Cached[string, any]{
		Cache: units.NewCache[string, any]("", nil, freq, ttl, units, opt...),
	}
	c.StartKeeping()
	return
}
func NewNumberCache(freq, ttl time.Duration, units units.MeasureUnit, opt ...units.Option) (c *Cached[float64, any]) {
	c = &Cached[float64, any]{
		Cache: units.NewCache[float64, any](0, nil, freq, ttl, units, opt...),
	}
	c.StartKeeping()
	return
}
