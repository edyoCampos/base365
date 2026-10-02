package crypto

import (
	"strings"
	"testing"
)

func TestGenerateAPIKey_Format(t *testing.T) {
	raw, hash, prefix, err := GenerateAPIKey()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(raw, "goclaw_") {
		t.Fatalf("api key must start with the product prefix, got %q", raw)
	}
	if len(raw) != len("goclaw_")+32 {
		t.Fatalf("api key must be prefix + 32 hex chars, got %q", raw)
	}
	if hash != HashAPIKey(raw) {
		t.Fatal("hash must match HashAPIKey(raw)")
	}
	if len(prefix) != 8 || !strings.HasPrefix(strings.TrimPrefix(raw, "goclaw_"), prefix) {
		t.Fatalf("display prefix %q must be the first 8 hex chars of the random part", prefix)
	}
}
