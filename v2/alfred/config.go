package alfred

import (
	"errors"
	"io"
	"os"

	"github.com/o98k-ok/lazy/v2/mac"
)

const InfoFile = "info.plist"

func FlowVariablesWithReader(reader io.ReadSeeker) (map[string]string, error) {
	var res map[string]string
	variables, err := mac.DefaultsRead(reader, []string{"variables"})
	if err != nil {
		return res, err
	}

	val, ok := variables.(map[string]interface{})
	if !ok {
		return res, errors.New("variables format error")
	}

	res = make(map[string]string, len(val))
	for k, v := range val {
		if vv, ok := v.(string); !ok {
			return res, errors.New("variables type error")
		} else {
			res[k] = vv
		}
	}
	return res, nil
}

func FlowVariables() (map[string]string, error) {
	var res map[string]string
	fi1e, err := os.Open(InfoFile)
	if err != nil {
		return res, err
	}

	defer fi1e.Close()
	return FlowVariablesWithReader(fi1e)
}
