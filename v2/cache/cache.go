package cache

import "time"

type Cache[T any] interface {
	Load(func(cached T, etime time.Time) (T, error)) (T, error)
	AsyncLoad(func(cached T, etime time.Time) (T, error)) (T, chan T, error)
}
