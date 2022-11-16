package database

import (
	"DJ-Blog/pkg/config"
	"testing"
)

func TestInitDB(t *testing.T) {
	config.Init("../../config")
	if !Init() {
		t.Error("Database init failed")
	}
}
