package main

import (
	"DJ-Blog/boot"
	"testing"
)

func TestMain(m *testing.M) {
	boot.Initialize()
	m.Run()
}
