package cache

type Cache[T any] interface {
	Load(func() (T, error)) (T, error)
	AsyncLoad(func() (T, error)) (T, chan T, error)
}
