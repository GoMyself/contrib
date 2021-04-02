package helper

import (
	"github.com/wI2L/jettison"
)


func JsonMarshal(v interface{}) ([]byte, error) {

	return jettison.Marshal(v)
}