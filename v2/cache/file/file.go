package file

import (
	"encoding/json"
	"io"
	"sync"
	"time"
)

// Cache cache T into file via dump
type Cache[T any] struct {
	mtx    sync.Mutex
	writer io.ReadWriter
	expire time.Duration
	codec  Codec
}

func NewFileCache[T any](rw io.ReadWriter, expire time.Duration) *Cache[T] {
	return &Cache[T]{
		mtx:    sync.Mutex{},
		expire: expire,
		writer: rw,
		codec:  JsonCodec{json.NewEncoder(rw), json.NewDecoder(rw)},
	}
}

func (f *Cache[T]) WithOption(codec Codec) *Cache[T] {
	f.codec = codec
	return f
}

func (f *Cache[T]) fetchAndDump(fn func() (T, error)) (res T, err error) {
	res, err = fn()
	if err != nil {
		return
	}

	f.mtx.Lock()
	err = f.codec.Encode(Persistence[T]{time.Now(), res})
	f.mtx.Unlock()
	return
}

// Load from cache file when cache is available
//      sync update and load when cache is unavailable
func (f *Cache[T]) Load(fn func() (T, error)) (T, error) {
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
	return f.fetchAndDump(fn)
}

// AsyncLoad from cache file when cache is available
//      async update and load when cache is unavailable
func (f *Cache[T]) AsyncLoad(fn func() (T, error)) (cached T, realtime chan T, err error) {
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
		cached, err = f.fetchAndDump(fn)
		realtime <- cached
		return
	}

	// cache is expired
	// fetch and cache it
	go func() {
		realData, _ := f.fetchAndDump(fn)
		realtime <- realData
	}()
	return persistenceCache.Content, realtime, err
}
