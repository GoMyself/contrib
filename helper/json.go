package helper

import (
	"fmt"
	"github.com/goccy/go-json"
)

func JsonMarshal(v interface{}) ([]byte, error) {
	fmt.Println("JsonMarshal", v)
	return json.Marshal(v)
}

func JsonUnmarshal(data []byte, v interface{}) error {
	fmt.Println("JsonUnmarshal", data)
	return json.Unmarshal(data, v)
}
