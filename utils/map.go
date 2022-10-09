package util

import (
	"encoding/json"
	"reflect"
)

type M map[string]interface{}

func (m M) MarshalJSON() ([]byte, error) {
	copyM := map[string]interface{}{}
	for k, v := range m {
		copyM[k] = v
		switch ref := reflect.ValueOf(v); ref.Kind() {
		case reflect.Slice:
			if ref.IsNil() {
				copyM[k] = json.RawMessage(`[]`)
			}
		}
	}
	return json.Marshal(copyM)
}
