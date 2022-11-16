package config

import "testing"

func TestInitConfig(t *testing.T) {
	InitConfig("../../config")
	port := LoadInt("server.port")
	if port == 0 {
		t.Error("load config error")
	}
}
