package helper

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func JsonMarshal(v interface{}) ([]byte, error) {
	fmt.Println("JsonMarshal", v)
	return json.Marshal(v)
}

func JsonUnmarshal(data []byte, v interface{}) error {
	fmt.Println("JsonUnmarshal", data)
	return json.Unmarshal(data, v)
}
