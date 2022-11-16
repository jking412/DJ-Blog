package config

import "testing"

func TestInitConfig(t *testing.T) {
	Init("../../config")
	port := LoadString("server.port")
	if port == "" {
		t.Error("load config error")
	}
}
