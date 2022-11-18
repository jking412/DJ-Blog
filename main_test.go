package main

import (
	"encoding/json"
	"testing"
)

type A struct {
	AbCdEf int
}

func TestTemp(t *testing.T) {
	a := A{
		AbCdEf: 1,
	}
	marshal, _ := json.Marshal(a)
	t.Log(string(marshal))
}
