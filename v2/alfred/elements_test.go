package alfred

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

func TestErrItems(t *testing.T) {
	t.Run("test items cases", func(t *testing.T) {
		title := "test cases"
		err := errors.New("test failed err")
		expect := fmt.Sprintf("{\"items\":[{\"title\":\"%s\",\"subtitle\":\"%s\"}]}", title, err.Error())

		logBuf := &bytes.Buffer{}
		expectLog := fmt.Sprintf("%s err %v", title, err)

		res := errItemsWithLog(title, err, logBuf)
		if res.Encode() != expect {
			t.Errorf("test failed, expect %v got %v", expect, res.Encode())
		}

		if logBuf.String() != expectLog {
			t.Errorf("test failed, expect %v got %v", expectLog, logBuf.String())
		}
	})
}
