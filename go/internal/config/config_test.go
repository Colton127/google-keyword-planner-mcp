package config

import "testing"

func TestIsCompleteWithoutDeveloperToken(t *testing.T) {
	cfg := Config{
		ClientID: "id", ClientSecret: "secret", RefreshToken: "refresh", CustomerID: "123",
	}
	if !cfg.IsComplete() {
		t.Fatal("OAuth credentials and customer ID should suffice")
	}
	cfg.RefreshToken = ""
	if cfg.IsComplete() {
		t.Fatal("refresh token is required")
	}
}
