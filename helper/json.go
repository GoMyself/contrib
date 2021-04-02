package helper

import (
	"github.com/wI2L/jettison"
	jsoniter "github.com/json-iterator/go"
)

var cjson = jsoniter.ConfigCompatibleWithStandardLibrary

func JsonMarshal(v interface{}) ([]byte, error) {

	return jettison.Marshal(v)
}

func JsonUnmarshal(data []byte, v interface{}) error {
	return cjson.Unmarshal(data, v)
}
