package file

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func genData(t *testing.T, tt time.Time, input string) string {
	d, err := json.Marshal(Persistence[string]{
		Time:    tt,
		Content: input,
	})
	assert.NoError(t, err)
	return string(d)
}

func createTempFile(t *testing.T, match string, content string) (*os.File, func()) {
	fd, err := os.CreateTemp("", match)
	assert.NoError(t, err)

	fd.Write([]byte(content))
	fd.Seek(0, 0)

	cleaner := func() {
		fd.Close()
		os.Remove(fd.Name())
	}
	return fd, cleaner
}

func TestSyncCache(t *testing.T) {
	t.Run("empty file", func(t *testing.T) {
		fd, cleaner := createTempFile(t, "sync_empty_file", "")
		defer cleaner()

		fCache := NewFileCache[string](fd, time.Second)
		got, err := fCache.Load(func(cached string, etime time.Time) (string, error) {
			return "hello, world", nil
		})
		assert.NoError(t, err)
		assert.Equal(t, "hello, world", got)

		gotd, err := ioutil.ReadFile(fd.Name())
		assert.NoError(t, err)
		assert.NotEmpty(t, gotd)
	})

	t.Run("with cache", func(t *testing.T) {
		var (
			cached   = "cached hello, world"
			realtime = "new hello, world"
		)

		d := genData(t, time.Now(), cached)
		fd, cleaner := createTempFile(t, "sync_with_cache", d)
		defer cleaner()

		// expire time is so long, so using cached result
		fCache := NewFileCache[string](fd, time.Hour)
		got, err := fCache.Load(func(cached string, etime time.Time) (string, error) {
			return realtime, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, cached, got)

		gotd, err := ioutil.ReadFile(fd.Name())
		assert.NoError(t, err)
		assert.Equal(t, d, string(gotd))
	})

	t.Run("with expire cache", func(t *testing.T) {
		input := "cache hello, world"
		d := genData(t, time.Now().Add(-time.Minute), input)
		fd, cleaner := createTempFile(t, "sync_with_cache", d)
		defer cleaner()

		// expire time is so short, so using new fetch data
		fCache := NewFileCache[string](fd, time.Second)
		got, err := fCache.Load(func(cached string, etime time.Time) (string, error) {
			return cached + cached, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, input+input, got)

		gotd, err := ioutil.ReadFile(fd.Name())
		assert.NoError(t, err)
		assert.NotEqual(t, d, string(gotd))
	})
}

func TestAsyncLoad(t *testing.T) {
	t.Run("empty file", func(t *testing.T) {
		fd, cleaner := createTempFile(t, "sync_empty_file", "")
		defer cleaner()

		fCache := NewFileCache[string](fd, time.Second)
		old, newd, err := fCache.AsyncLoad(func(cached string, etime time.Time) (string, error) {
			return "hello, world", nil
		})

		assert.NoError(t, err)
		assert.Equal(t, "hello, world", old)
		assert.Equal(t, "hello, world", <-newd)

		gotd, err := ioutil.ReadFile(fd.Name())
		assert.NoError(t, err)
		assert.NotEmpty(t, gotd)
	})

	t.Run("with cache", func(t *testing.T) {
		// {"Time":"2022-05-14T14:00:58.483194+08:00","Content":"hello, world"}
		var (
			cached   = "cached hello, world"
			realtime = "new hello, world"
		)

		d := genData(t, time.Now(), cached)
		fd, cleaner := createTempFile(t, "sync_with_cache", d)
		defer cleaner()

		// expire time is so long, so using cached result
		fCache := NewFileCache[string](fd, time.Hour)
		old, newd, err := fCache.AsyncLoad(func(cached string, etime time.Time) (string, error) {
			return realtime, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, cached, old)
		assert.Equal(t, cached, <-newd)

		gotd, err := ioutil.ReadFile(fd.Name())
		assert.NoError(t, err)
		assert.Equal(t, d, string(gotd))
	})

	t.Run("with expire cache", func(t *testing.T) {
		// {"Time":"2022-05-14T14:00:58.483194+08:00","Content":"hello, world"}
		var (
			cached   = "cached hello, world"
			realtime = "new hello, world"
		)
		d := genData(t, time.Now().Add(-time.Minute), "cached hello, world")
		fd, cleaner := createTempFile(t, "sync_with_cache", d)
		defer cleaner()

		// expire time is so short, so using new fetch data
		fCache := NewFileCache[string](fd, time.Second)
		old, newd, err := fCache.AsyncLoad(func(cached string, etime time.Time) (string, error) {
			return realtime, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, cached, old)
		assert.Equal(t, realtime, <-newd)

		gotd, err := ioutil.ReadFile(fd.Name())
		assert.NoError(t, err)
		assert.NotEqual(t, d, string(gotd))
	})
}
