package main

import (
	"DJ-Blog/pkg/helper"
	"testing"
)

func TestTemp(t *testing.T) {
	t.Log(helper.GetUTF8StringLength(("我是")))
}
