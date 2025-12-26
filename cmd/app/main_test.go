package main

import (
	"os"
	"testing"
)

func TestGetEnvWithFallback(t *testing.T) {
	os.Setenv("TEST_KEY", "custom")
	defer os.Unsetenv("TEST_KEY")

	value := getEnv("TEST_KEY", "default")
	if value != "custom" {
		t.Fatalf("expected fallback value, got %s", value)
	}

}

func TestServerStartStop(t *testing.T) {
	server := &Server{Name: "test"}

	server.Start()
	if !server.Active {
		t.Fatalf("expected server to be active after start")
	}
	server.Stop()
	if server.Active {
		t.Fatalf("expected server to be inactive after stop")
	}
}
