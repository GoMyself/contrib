package helper

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/wI2L/jettison"
)

var cjson = jsoniter.ConfigCompatibleWithStandardLibrary

func JsonMarshal(v interface{}) ([]byte, error) {

	return jettison.Marshal(v)
}

func JsonUnmarshal(data []byte, v interface{}) error {
	return cjson.Unmarshal(data, v)
}
