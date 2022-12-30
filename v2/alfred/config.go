package alfred

import (
	"errors"
	"io"
	"os"
	"strconv"

	"github.com/o98k-ok/lazy/v2/mac"
)

const InfoFile = "info.plist"

type Envs map[string]string

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

func GetFlowEnv() (Envs, error) {
	return FlowVariables()
}

func (e Envs) GetAsInt(key string, def int) int {
	val, ok := e[key]
	if !ok {
		return def
	}

	res, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return def
	}
	return int(res)
}

func (e Envs) GetAsString(key string, def string) string {
	val, ok := e[key]
	if !ok {
		return def
	}

	if len(val) != 0 {
		return val
	}
	return def
}
