package config

import "testing"

func TestInitConfig(t *testing.T) {
	InitConfig("../../config")
	port := LoadString("server.port")
	if port == "" {
		t.Error("load config error")
	}
}
