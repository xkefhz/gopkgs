package internal

import (
	"bytes"
	"encoding/json"
)

func EqualJSON(a, b interface{}) bool {
	bufa, _ := json.Marshal(a)
	bufb, _ := json.Marshal(b)
	return bytes.Equal(bufa, bufb)
}
