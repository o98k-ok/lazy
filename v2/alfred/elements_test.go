package alfred

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrItems(t *testing.T) {
	t.Run("test items cases", func(t *testing.T) {
		title := "test cases"
		err := errors.New("test failed err")
		expect := fmt.Sprintf("{\"items\":[{\"title\":\"%s\",\"subtitle\":\"%s\"}]}", title, err.Error())

		logBuf := &bytes.Buffer{}
		expectLog := fmt.Sprintf("%s err %v\n", title, err)

		res := errItemsWithLog(title, err, logBuf)
		assert.Equal(t, expect, res.Encode())
		assert.Equal(t, expectLog, logBuf.String())
	})
}
