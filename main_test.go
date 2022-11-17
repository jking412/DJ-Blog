package main

import (
	"testing"
	"time"
)

func TestTemp(t *testing.T) {
	t.Log(time.Now().Format("2006-01-02 15:04:05.000000000"))
}
