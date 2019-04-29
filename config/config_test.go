package config

import (
	"encoding/json"
	"testing"
)

func testEqual(a, b interface{}) bool {
	a1, _ := json.Marshal(a)
	b1, _ := json.Marshal(b)
	return string(a1) == string(b1)
}

func TestLoadConfig(t *testing.T) {
	env1 := &Env{
		Sign:       "420e57b017066b44e05ea1577f6e2e12",
		ProductKey: "helloworld!",
		ServerList: []server{
			server{Name: "router", Addr: "127.0.0.1:9003"},
		},
	}
	env2 := &Env{}
	LoadFile("config.xml", env2)
	if !testEqual(env1, env2) {
		t.Error("not equal")
	}
}

func TestParseArgs(t *testing.T) {
	samples := [][]string{
		{"1", "-test=1", "-test2=2"},
		{"", "-test2=1", "-test3=2"},
		{"1", "-test", "1", "-test3=2"},
		{"1", "-test2=1", "-test", "1", "-test3=2"},
		{"abcde1", "-test2", "1", "-test", "abcde1", "-test3=2"},
	}
	for _, sample := range samples {
		v := ParseCmdArgs(sample[1:], "test", "")
		if v != sample[0] {
			t.Errorf("parse %v result: %s", sample, v)
		}
	}
}
