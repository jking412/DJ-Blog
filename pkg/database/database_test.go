package database

import (
	"DJ-Blog/pkg/config"
	"testing"
)

func TestInitDB(t *testing.T) {
	config.InitConfig("../../config")
	err := InitDB()
	if err != nil {
		t.Error(err)
	}
}
