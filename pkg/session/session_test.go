package session

import (
	"DJ-Blog/pkg/config"
	"testing"
)

func TestInitSession(t *testing.T) {
	config.InitConfig("../../config")
	InitSession()
}
