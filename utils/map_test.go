package utils

import (
	"encoding/json"
	"github.com/xkefhz/gopkgs/internal"
	"testing"
)

type mStruct struct {
	N int
}

func TestM(t *testing.T) {
	m1 := M{
		"Slice":     []int{1, 2},
		"NilSlice":  ([]int)(nil),
		"Struct":    mStruct{N: 1},
		"NilStruct": (*mStruct)(nil),
		"Nil":       nil,
		"String":    "StringA",
	}
	m2 := map[string]interface{}{
		"Slice":     []int{1, 2},
		"NilSlice":  []int{},
		"Struct":    mStruct{N: 1},
		"NilStruct": (*mStruct)(nil),
		"Nil":       nil,
		"String":    "StringA",
	}
	if !internal.EqualJSON(m1, m2) {
		buf1, _ := json.Marshal(m1)
		buf2, _ := json.Marshal(m2)
		t.Errorf("check M.MarshalJSON m1:%s m2:%s", buf1, buf2)
	}
}
