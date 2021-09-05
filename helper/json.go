package helper

import (
	"github.com/goccy/go-json"
)


func JsonMarshal(v interface{}) ([]byte, error) {

	return json.Marshal(v)
}

func JsonUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
