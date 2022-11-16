package database

import (
	"DJ-Blog/pkg/config"
	"testing"
)

func TestInitDB(t *testing.T) {
	config.InitConfig("../../config")
	if !InitDB() {
		t.Error("Database init failed")
	}
}
