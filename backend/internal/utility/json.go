package utility

import (
	"encoding/json"
	"errors"
	"strings"
)

func MarshalToJSON(wrapper []string, i ...interface{}) (string, error) {

	if len(wrapper) != len(i) {
		return "", errors.New("wrapper length does not match interfaces length")
	}
	js := "{"
	for k, v := range i {
		s, err := json.Marshal(v)
		if err != nil {
			return "", err
		}
		js = js + `"` + wrapper[k] + `":` + string(s) + ","
	}
	js = strings.TrimRight(js, ",")
	return js + "}", nil
}
