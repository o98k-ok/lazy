package file

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

// Cache cache T into file via dump
type Cache[T any] struct {
	mtx    sync.Mutex
	fd     *os.File
	expire time.Duration
	codec  Codec
}

func NewFileCache[T any](fd *os.File, expire time.Duration) *Cache[T] {
	return &Cache[T]{
		mtx:    sync.Mutex{},
		expire: expire,
		fd:     fd,
		codec:  JsonCodec{json.NewEncoder(fd), json.NewDecoder(fd)},
	}
}

func (f *Cache[T]) WithOption(codec Codec) *Cache[T] {
	f.codec = codec
	return f
}

func (f *Cache[T]) fetchAndDump(fn func(cached T, etime time.Time) (T, error), cached Persistence[T]) (res T, err error) {
	res, err = fn(cached.Content, cached.Time)
	if err != nil {
		return
	}

	f.mtx.Lock()
	f.fd.Truncate(0)
	f.fd.Seek(0, 0)
	err = f.codec.Encode(Persistence[T]{time.Now(), res})
	f.mtx.Unlock()
	return
}

// Load from cache file when cache is available
//      sync update and load when cache is unavailable
func (f *Cache[T]) Load(fn func(cached T, etime time.Time) (T, error)) (T, error) {
	var cached Persistence[T]
	f.mtx.Lock()
	err := f.codec.Decode(&cached)
	f.mtx.Unlock()

	// get available cache, just return it
	if err == nil && !cached.Time.Add(f.expire).Before(time.Now()) {
		return cached.Content, nil
	}

	// cache is expired or not available
	// fetch and cache it
	return f.fetchAndDump(fn, cached)
}

// AsyncLoad from cache file when cache is available
//      async update and load when cache is unavailable
func (f *Cache[T]) AsyncLoad(fn func(cached T, etime time.Time) (T, error)) (cached T, realtime chan T, err error) {
	realtime = make(chan T, 1)
	var persistenceCache Persistence[T]
	f.mtx.Lock()
	err = f.codec.Decode(&persistenceCache)
	f.mtx.Unlock()

	// get available cache, just return it
	if err == nil && !persistenceCache.Time.Add(f.expire).Before(time.Now()) {
		cached = persistenceCache.Content
		realtime <- persistenceCache.Content
		return
	}

	// cache is not available, fetch and cache it
	if err != nil {
		cached, err = f.fetchAndDump(fn, persistenceCache)
		realtime <- cached
		return
	}

	// cache is expired
	// fetch and cache it
	go func() {
		realData, _ := f.fetchAndDump(fn, persistenceCache)
		realtime <- realData
	}()
	return persistenceCache.Content, realtime, err
}
