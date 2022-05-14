package file

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSyncCache(t *testing.T) {
	t.Run("empty file", func(t *testing.T) {
		buff := &bytes.Buffer{}
		fCache := NewFileCache[string](buff, time.Second)
		got, err := fCache.Load(func() (string, error) {
			return "hello, world", nil
		})
		assert.NoError(t, err)
		assert.Equal(t, "hello, world", got)

		var per Persistence[string]
		err = json.Unmarshal(buff.Bytes(), &per)
		assert.NoError(t, err)
		assert.Equal(t, "hello, world", per.Content)
	})

	t.Run("with cache", func(t *testing.T) {
		// {"Time":"2022-05-14T14:00:58.483194+08:00","Content":"hello, world"}
		var (
			cached   = "cached hello, world"
			realtime = "new hello, world"
		)
		d, err := json.Marshal(Persistence[string]{
			Time:    time.Now(),
			Content: cached,
		})
		assert.NoError(t, err)
		buff := bytes.NewBuffer(d)

		// expire time is so long, so using cached result
		fCache := NewFileCache[string](buff, time.Hour)
		got, err := fCache.Load(func() (string, error) {
			return realtime, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, cached, got)

		// buff has read over, so check empty buff
		assert.Empty(t, buff.Bytes())
	})

	t.Run("with expire cache", func(t *testing.T) {
		// {"Time":"2022-05-14T14:00:58.483194+08:00","Content":"hello, world"}
		var (
			cached   = "cached hello, world"
			realtime = "new hello, world"
		)
		d, err := json.Marshal(Persistence[string]{
			Time:    time.Now().Add(-time.Minute),
			Content: cached,
		})
		assert.NoError(t, err)
		buff := bytes.NewBuffer(d)

		// expire time is so short, so using new fetch data
		fCache := NewFileCache[string](buff, time.Second)
		got, err := fCache.Load(func() (string, error) {
			return realtime, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, realtime, got)

		// check cache updated
		assert.NotEqual(t, d, buff.Bytes())
	})
}

func TestAsyncLoad(t *testing.T) {
	t.Run("empty file", func(t *testing.T) {
		buff := &bytes.Buffer{}
		fCache := NewFileCache[string](buff, time.Second)
		old, newd, err := fCache.AsyncLoad(func() (string, error) {
			return "hello, world", nil
		})

		assert.NoError(t, err)
		assert.Equal(t, "hello, world", old)
		assert.Equal(t, "hello, world", <-newd)

		var per Persistence[string]
		err = json.Unmarshal(buff.Bytes(), &per)
		assert.NoError(t, err)
		assert.Equal(t, "hello, world", per.Content)
	})

	t.Run("with cache", func(t *testing.T) {
		// {"Time":"2022-05-14T14:00:58.483194+08:00","Content":"hello, world"}
		var (
			cached   = "cached hello, world"
			realtime = "new hello, world"
		)
		d, err := json.Marshal(Persistence[string]{
			Time:    time.Now(),
			Content: cached,
		})
		assert.NoError(t, err)
		buff := bytes.NewBuffer(d)

		// expire time is so long, so using cached result
		fCache := NewFileCache[string](buff, time.Hour)
		old, newd, err := fCache.AsyncLoad(func() (string, error) {
			return realtime, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, cached, old)
		assert.Equal(t, cached, <-newd)

		// buff has read over, so check empty buff
		assert.Empty(t, buff.Bytes())
	})

	t.Run("with expire cache", func(t *testing.T) {
		// {"Time":"2022-05-14T14:00:58.483194+08:00","Content":"hello, world"}
		var (
			cached   = "cached hello, world"
			realtime = "new hello, world"
		)
		d, err := json.Marshal(Persistence[string]{
			Time:    time.Now().Add(-time.Minute),
			Content: cached,
		})
		assert.NoError(t, err)
		buff := bytes.NewBuffer(d)

		// expire time is so short, so using new fetch data
		fCache := NewFileCache[string](buff, time.Second)
		old, newd, err := fCache.AsyncLoad(func() (string, error) {
			return realtime, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, cached, old)
		assert.Equal(t, realtime, <-newd)

		// check cache updated
		assert.NotEqual(t, d, buff.Bytes())
	})
}
