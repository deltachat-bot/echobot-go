package main

import (
	"testing"

	"github.com/deltachat/deltachat-rpc-client-go/acfactory"
)

func TestMain(m *testing.M) {
	cfg := map[string]string{
		"mail_server":   "localhost",
		"send_server":   "localhost",
		"mail_port":     "3143",
		"send_port":     "3025",
		"mail_security": "3",
		"send_security": "3",
	}
	acfactory.TearUp(cfg)
	defer acfactory.TearDown()
	m.Run()
}
